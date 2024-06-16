package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/mail"
	"strings"
	"time"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/Jaynxe/xie-blog/utils"
	"github.com/Jaynxe/xie-blog/utils/email"
	"github.com/Jaynxe/xie-blog/utils/errhandle"
	"github.com/Jaynxe/xie-blog/utils/pwd"
	"github.com/Jaynxe/xie-blog/utils/random"
	"github.com/Jaynxe/xie-blog/utils/snowflake"
	"github.com/Jaynxe/xie-blog/utils/token"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	RefreshTokenExpired = 24 * time.Hour * 7
	AccessTokenExpired  = 2 * time.Hour
)

// IsValidSession 是否登录 godoc
// @Summary 是否登录
// @Schemes
// @Description 是否登录
// @Tags auth
// @Accept json
// @Produce json
// @Param   Authorization  header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[model.GetUserResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /isvalid [get]
func (a *Auth) IsValidSession(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	prefix := "Bearer "
	tk := ""

	if auth != "" && strings.HasPrefix(auth, prefix) {
		tk = auth[len(prefix):]
	}

	if tk == "" {
		model.Throw(c, errhandle.TokenError)
		return
	}

	info, ok := token.TK.Verify(context.Background(), tk)
	if !ok {
		model.Throw(c, errhandle.PermissionDenied)
		return
	}

	var usr model.User
	err := global.GVB_DB.Table("users").
		Where("id = ?", info.UserID).
		First(&usr).Error

	if err != nil {
		model.ThrowError(c, err)
		return
	}
	var getResponse model.GetUserResponse

	utils.IgnoreStructCopy(&getResponse, &usr, "")

	model.OK[model.GetUserResponse](c, getResponse)
}

// UserLogin 用户名密码登录 godoc
// @Summary 用户名密码登录
// @Schemes
// @Description 用户名密码登录
// @Tags auth
// @Accept json
// @Produce json
// @Param   userInfo  body    model.UserLoginRequest  true   "用户名, 密码"
// @Success 200 {object} model.CommonResponse[model.TokenResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /login [post]
func (a *Auth) UserLogin(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	var loginReq model.UserLoginRequest
	err = json.Unmarshal(b, &loginReq)
	if err != nil {
		model.ThrowError(c, err)
		return
	}

	tx := utils.BuildLoginSQL(global.GVB_DB.Table("users"), &loginReq)
	if tx == nil {
		model.Throw(c, errhandle.ParamsError)
		return
	}

	var user model.User
	err = tx.First(&user).Error
	if err != nil {
		model.Throw(c, errhandle.UserNonExists)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(loginReq.Password),
	)

	if err != nil {
		model.Throw(c, errhandle.PasswordInvalid)
		return
	}

	accessToken, err := token.TK.Token(user.ID, user.Role, user.Name, AccessTokenExpired)
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	refreshToken, err := token.TK.Token(user.ID, user.Role, user.Name, RefreshTokenExpired)
	if err != nil {
		model.ThrowError(c, err)
		return
	}

	model.OK[model.TokenResponse](c, model.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Scope:        user.Role,
		ExpiredAt:    time.Now().Add(AccessTokenExpired).Unix(),
	})
}

// UserLoginRefresh 刷新登录令牌 godoc
// @Summary 刷新登录令牌
// @Schemes
// @Description 刷新登录令牌
// @Tags auth
// @Accept json
// @Produce json
// @Param   Authorization    header    string    true  "用户Refresh Token"
// @Success 200 {object} model.CommonResponse[model.TokenResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /refresh [post]
func (a *Auth) UserLoginRefresh(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	const prefix = "Bearer "
	var tk string

	if auth != "" && strings.HasPrefix(auth, prefix) {
		tk = auth[len(prefix):]
	}

	if tk == "" {
		model.Throw(c, errhandle.TokenError)
		return
	}

	userinfo, ok := token.TK.Verify(context.Background(), tk)
	if !ok {
		model.Throw(c, errhandle.PermissionDenied)
		return
	}

	accessToken, err := token.TK.Token(userinfo.UserID, userinfo.Role, userinfo.Name, AccessTokenExpired)
	if err != nil {
		model.ThrowError(c, err)
		return
	}

	model.OK[model.TokenResponse](c, model.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: tk, //继续返回刷新令牌
		Scope:        userinfo.Role,
		ExpiredAt:    time.Now().Add(AccessTokenExpired).Unix(),
	})
}

// UserRegister 注册普通用户 godoc
// @Summary 注册普通用户
// @Schemes
// @Description 注册普通用户
// @Tags auth
// @Accept json
// @Produce json
// @Param   registerInfo  body   model.RegisterUserRequest  true  "用户注册信息"
// @Success 200 {object} model.CommonResponse[string]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /register [post]
func (a *Auth) UserRegister(c *gin.Context) {
	var req model.RegisterUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	session := sessions.Default(c)
	if req.VerificationCode == nil {
		// 第一次调用这个接口，生成四位验证码
		code := random.VerifyCode(4)

		// 把验证码写入session
		session.Set("verify_code_register", code)
		session.Set("email_register", req.Email)
		err := session.Save()
		if err != nil {
			global.GVB_LOGGER.Error("保存session失败", err)
			model.ThrowError(c, err)
			return
		}
		// 发送验证码到用户邮箱
		err = email.NewVerificationCodeApi().Send(req.Email, fmt.Sprintf("您的用户注册验证码是 [ %s ]", code))
		if err != nil {
			model.ThrowError(c, err)
			return
		}
		model.OK(c, "验证码已发送到邮箱")
		return
	}
	emailRegister := session.Get("email_register")
	if emailRegister == nil || req.Email != emailRegister {
		global.GVB_LOGGER.Error("两次邮箱不一致")
		model.Throw(c, errhandle.EmailIsDifferent)
		return
	}
	// 第二次调用这个接口进行验证码验证和密码重置
	code := session.Get("verify_code_register")
	global.GVB_LOGGER.Info("获取到的验证码: ", code)
	if code == nil {
		global.GVB_LOGGER.Error("验证码已过期或未设置")
		model.Throw(c, errhandle.VerifyCodeError)
		return
	}
	if *req.VerificationCode != code {
		global.GVB_LOGGER.Error("验证码错误")
		model.Throw(c, errhandle.VerifyCodeError)
		return
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		model.Throw(c, errhandle.EmailFormatError)
		return
	}
	if req.Sex != "男" && req.Sex != "女" {
		model.Throw(c, errhandle.SexError)
		return
	}
	if !utils.IsValidPassword(req.Password) {
		model.Throw(c, errhandle.PasswordFormatError)
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	var found model.User
	col := global.GVB_DB.Table("users").FirstOrCreate(&found, model.User{
		ID:       snowflake.ID(),
		Role:     "user",
		Name:     req.Name,
		NickName: req.NickName,
		Sex:      req.Sex,
		Password: string(hashed),
		Email:    req.Email,
		Avatar:   req.Avatar,
		IP:       c.ClientIP(),
	})
	// this shouldn't happen
	if col.RowsAffected == 0 {
		model.Throw(c, errhandle.InnerError)
		return
	}

	model.OK(c, "注册成功")
}

// GetAllArticles 获取所有文章 godoc
// @Summary 获取所有文章
// @Schemes
// @Description 获取所有文章
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} model.CommonResponse[model.GetUserResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /getallArticles [get]
func (a *Auth) GetAllArticles(c *gin.Context) {
	var allArticles []model.Article

	if err := global.GVB_DB.Preload("Comments").Preload("Tags").Find(&allArticles).Error; err != nil {
		model.ThrowError(c, err)
		return
	}
	model.OK[any](c, allArticles)
}

// GetAllMenus 获取所有菜单 godoc
// @Summary 获取所有菜单
// @Schemes
// @Description 获取所有菜单
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} model.CommonResponse[[]model.MenuItem]
// @Failure 400  {object} model.CommonResponse[any]
// @Router	/getAllMenus [get]
func (a *Auth) GetAllMenus(c *gin.Context) {
	var ml []model.MenuItem
	err := global.GVB_DB.Find(&ml).Error
	if err != nil {
		global.GVB_LOGGER.Error("菜单查询失败")
		model.ThrowError(c, err)
		return
	}
	model.OKWithMsg(c, ml, "菜单查询成功")
}

// LoginWithEmail 邮箱登录 godoc
// @Summary 邮箱登录
// @Schemes
// @Description 邮箱登录
// @Tags auth
// @Accept json
// @Produce json
// @Param  LoginWithEmailRequest   body    model.LoginWithEmailRequest  true   "邮箱, 密码, 验证码"
// @Success 200 {object} model.CommonResponse[model.TokenResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /loginWithEmail [post]
func (a *Auth) LoginWithEmail(c *gin.Context) {
	var ber model.LoginWithEmailRequest
	err := c.ShouldBindJSON(&ber)
	if err != nil {
		model.ThrowBindError(c, &ber, err)
		return
	}
	// 是否存在该用户
	var user model.User
	isExist := global.GVB_DB.First(&user, "email = ?", ber.Email).RowsAffected
	if isExist == 0 {
		global.GVB_LOGGER.Error("邮箱不存在")
		model.Throw(c, errhandle.UserNonExists)
		return
	}
	session := sessions.Default(c)
	if ber.VerificationCode == nil {
		// 第一次调用这个接口，后台生成四位验证码
		code := random.VerifyCode(4)

		// 把验证码写入session
		session.Set("verify_code_login", code)
		session.Set("email", ber.Email)
		err := session.Save()
		if err != nil {
			global.GVB_LOGGER.Error(err)
			model.ThrowError(c, err)
			return
		}

		// 发送到要登陆用户的邮箱
		err = email.NewVerificationCodeApi().Send(ber.Email, fmt.Sprintf("您的邮箱登录验证码是 [ %s ]", code))
		if err != nil {
			model.ThrowError(c, err)
			return
		}
		model.OK(c, "验证码获取成功")
		return
	}
	// 第二次调用这个接口进行登录校验
	code := session.Get("verify_code_login")
	if *ber.VerificationCode != code {
		global.GVB_LOGGER.Error("验证码错误")
		model.Throw(c, errhandle.VerifyCodeError) //验证码错误
		return
	}
	emailLogin := session.Get("email")
	if ber.Email != emailLogin {
		global.GVB_LOGGER.Error("两次邮箱不一致")
		model.Throw(c, errhandle.EmailIsDifferent)
		return
	}
	// 发放token
	accessToken, err := token.TK.Token(user.ID, user.Role, user.Name, AccessTokenExpired)
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	refreshToken, err := token.TK.Token(user.ID, user.Role, user.Name, AccessTokenExpired)
	if err != nil {
		model.ThrowError(c, err)
		return
	}

	model.OK[model.TokenResponse](c, model.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Scope:        user.Role,
		ExpiredAt:    time.Now().Add(AccessTokenExpired).Unix(),
	})

}

// ResetPassword /* 重置密码 */
func (a *Auth) ResetPassword(c *gin.Context) {
	var req model.ResetPasswordRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		model.ThrowBindError(c, &req, err)
		return
	}

	// 检查用户是否存在
	var user model.User
	isExist := global.GVB_DB.First(&user, "email = ?", req.Email).RowsAffected
	if isExist == 0 {
		global.GVB_LOGGER.Error("邮箱不存在")
		model.Throw(c, errhandle.UserNonExists)
		return
	}

	session := sessions.Default(c)
	if req.VerificationCode == nil {
		// 第一次调用这个接口，生成四位验证码
		code := random.VerifyCode(4)

		// 把验证码写入session
		session.Set("verify_code_reset", code)
		session.Set("email_reset", req.Email)
		err := session.Save()
		if err != nil {
			global.GVB_LOGGER.Error("保存session失败", err)
			model.ThrowError(c, err)
			return
		}
		// 发送验证码到用户邮箱
		err = email.NewVerificationCodeApi().Send(req.Email, fmt.Sprintf("您的重置密码验证码是 [ %s ]", code))
		if err != nil {
			model.ThrowError(c, err)
			return
		}
		model.OK(c, "验证码已发送到邮箱")
		return
	}
	emailReset := session.Get("email_reset")
	if emailReset == nil || req.Email != emailReset {
		global.GVB_LOGGER.Error("两次邮箱不一致")
		model.Throw(c, errhandle.EmailIsDifferent)
		return
	}
	// 第二次调用这个接口进行验证码验证和密码重置
	code := session.Get("verify_code_reset")
	global.GVB_LOGGER.Info("获取到的验证码: ", code)
	if code == nil {
		global.GVB_LOGGER.Error("验证码已过期或未设置")
		model.Throw(c, errhandle.VerifyCodeError)
		return
	}
	if *req.VerificationCode != code {
		global.GVB_LOGGER.Error("验证码错误")
		model.Throw(c, errhandle.VerifyCodeError)
		return
	}
	// 重置密码
	if !utils.IsValidPassword(req.NewPassword) {
		model.Throw(c, errhandle.PasswordFormatError)
		return
	}
	user.Password = pwd.HashAndSalt(req.NewPassword)
	err = global.GVB_DB.Save(&user).Error
	if err != nil {
		global.GVB_LOGGER.Error("更新密码失败", err)
		model.ThrowError(c, err)
		return
	}

	// 清除session
	session.Delete("verify_code_reset")
	session.Delete("email_reset")
	err = session.Save()
	if err != nil {
		global.GVB_LOGGER.Error(err)
		model.ThrowError(c, err)
		return
	}

	model.OK(c, "密码重置成功")
}

// LoginWithQQ 待续....
