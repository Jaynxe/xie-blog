package admin

import (
	"context"
	"encoding/json"
	"errors"
	"net/mail"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/Jaynxe/xie-blog/utils"
	"github.com/Jaynxe/xie-blog/utils/errhandle"
	"github.com/Jaynxe/xie-blog/utils/pwd"
	"github.com/Jaynxe/xie-blog/utils/snowflake"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	cb = context.Background()
)

/* ------------用户注册-------- */
// 注册管理员 godoc
// @Summary 注册管理员
// @Schemes
// @Description 注册管理员
// @Tags admin
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   registerInfo  body   model.RegisterUserRequest  true  "用户注册信息"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/admin/new [post]
func (a *Admin) RegisterAdmin(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	var req model.RegisterUserRequest
	json.Unmarshal(b, &req)

	if _, err := mail.ParseAddress(req.Email); err != nil {
		model.Throw(c, errhandle.EmailFormatError)
		return
	}
	if req.Sex != "男" && req.Sex != "女" {
		model.Throw(c, errhandle.SexError)
		return
	}
	if !utils.IsValidPassword(req.Password) {
		model.Throw(c, errhandle.PasswordTooShort)
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	var found model.User
	col := global.GVB_DB.Table("users").FirstOrCreate(&found, model.User{
		ID:       snowflake.ID(),
		Role:     "admin",
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

	model.OK[any](c, nil)
}

/* ------------获取用户信息-------- */

// 分页获取用户 godoc
// @Summary 分页获取用户
// @Schemes
// @Description 分页获取用户
// @Tags admin
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   page    query   int  false  "页码"
// @Param   key    query   string  false  "搜索关键字"
// @Param   limit    query   int  false  "每页大小"
// @Param   sort    query   string  false  "排序规则"
// @Success 200 {object} model.CommonResponse[[]model.User]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/paginatedUsers [get]
func (a *Admin) GetPaginatedUsers(c *gin.Context) {
	var page model.PageRequest
	err := c.ShouldBindQuery(&page)
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	UserList := []model.User{}
	UserList, err = utils.ComList(UserList, utils.Option{PageRequest: page, Debug: false})
	if err != nil {
		global.GVB_LOGGER.Error(err.Error())
		model.ThrowError(c, err)
		return
	}
	model.OKWithMsg(c, UserList, "查询用户列表成功")
}

// 获取所有用户 godoc
// @Summary 获取所有用户
// @Schemes
// @Description 获取所有用户
// @Tags admin
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[[]model.User]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/getAllUsers [get]
func (a *Admin) GetAllUsers(c *gin.Context) {
	var UserList []model.User
	if err := global.GVB_DB.Find(&UserList).Error; err != nil {
		global.GVB_LOGGER.Error("获取用户列表失败")
		model.ThrowError(c, err)
		return
	}
	model.OKWithMsg(c, UserList, "查询成功")
}

/* ------------修改用户信息-------- */

// 修改指定管理员信息 godoc
// @Summary 修改指定管理员信息
// @Schemes
// @Description 修改指定管理员信息
// @Tags admin
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   ModifyUserRequest     body     model.ModifyUserRequest  true  "要修改的管理员的信息"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/modifyAdmin [patch]
func (t *Admin) ModifyAdmin(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	var req model.ModifyUserRequest
	json.Unmarshal(b, &req)

	var user model.User
	err = global.GVB_DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "admin").
		First(&user).Error
	if err != nil {
		model.Throw(c, errhandle.UserNonExists)
		return
	}

	utils.IgnoreStructCopy(&user, &req, "")

	global.GVB_DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "admin").
		Save(&user)

	model.OK[any](c, nil)
}

// 修改指定用户信息 godoc
// @Summary 修改指定用户信息
// @Schemes
// @Description 修改指定用户信息
// @Tags admin
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   ModifyUserRequest     body     model.ModifyUserRequest  true  "要修改的用户的信息"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/modifyUser [patch]
func (a *Admin) ModifyUser(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	var req model.ModifyUserRequest
	json.Unmarshal(b, &req)

	var user model.User
	err = global.GVB_DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "user").
		First(&user).Error
	if err != nil {
		model.Throw(c, errhandle.UserNonExists)
		return
	}

	utils.IgnoreStructCopy(&user, &req, "")

	global.GVB_DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "user").
		Save(&user)

	model.OK[any](c, nil)
}

// 修改指定管理员密码 godoc
// @Summary 修改指定管理员密码
// @Schemes
// @Description 修改指定管理员密码
// @Tags admin
// @Accept json
// @Produce json
// @Param   Authorization     header    string  true   "登录返回的Token"
// @Param   ModifyPasswordRequest    body   model.ModifyPasswordRequest  true  "id,新密码"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/modifyAdminPassword [patch]
func (t *Admin) ModifyAdminPassword(c *gin.Context) {
	info, ok := c.Get("info")
	if !ok {
		model.ThrowWithMsg(c, "获取上下文中的用户信息失败")
		return
	}
	userInfo := info.(*model.UserInfo)
	b, err := c.GetRawData()
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	var req model.ModifyPasswordRequest
	json.Unmarshal(b, &req)

	var u model.User
	err = global.GVB_DB.Take(&u, req.UserID).Error
	if err != nil {
		model.Throw(c, errhandle.UserNonExists)
		return
	}

	if !pwd.ComparePasswords(u.Password, req.OldPwd) {
		model.ThrowWithMsg(c, "旧密码和数据库的不一致")
		return
	}
	if req.NewPwb == "" {
		model.Throw(c, errhandle.ParamsError)
		return
	}
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPwb),
		bcrypt.DefaultCost,
	)
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	result := global.GVB_DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "admin").
		Update("password", string(hashed))
	if result.Error != nil {
		model.Throw(c, errhandle.UserNonExists)
		return
	}
	if result.RowsAffected == 0 {
		// 没有任何记录被更新，可能是因为用户不存在或角色不匹配
		global.GVB_LOGGER.Warn("用户不存在或角色不匹配")
		model.Throw(c, errhandle.UserNonExists)
		return
	}
	// 如果被修改密码的管理员刚好是当前在线的
	if req.UserID == userInfo.UserID {
		global.GVB_REDIS.Del(cb, userInfo.UUID)
	}
	model.OK[any](c, nil)
}

// 修改指定用户密码 godoc
// @Summary 修改指定用户密码
// @Schemes
// @Description 修改指定用户密码
// @Tags admin
// @Accept json
// @Produce json
// @Param   Authorization     header    string  true   "登录返回的Token"
// @Param   ModifyPasswordRequest    body   model.ModifyPasswordRequest  true  "id,新密码"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/modifyUserPassword [patch]
func (t *Admin) ModifyUserPassword(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	var req model.ModifyPasswordRequest
	json.Unmarshal(b, &req)

	var user model.User
	err = global.GVB_DB.Take(&user, req.UserID).Error
	if err != nil {
		model.Throw(c, errhandle.UserNonExists)
		return
	}

	if !pwd.ComparePasswords(user.Password, req.OldPwd) {
		model.ThrowWithMsg(c, "旧密码和数据库的不一致")
		return
	}

	if req.NewPwb == "" {
		model.Throw(c, errhandle.ParamsError)
		return
	}
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPwb),
		bcrypt.DefaultCost,
	)
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	result := global.GVB_DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "user").
		Update("password", string(hashed))
	if result.Error != nil {
		model.Throw(c, errhandle.UserNonExists)
		return
	}
	if result.RowsAffected == 0 {
		// 没有任何记录被更新，可能是因为用户不存在或角色不匹配
		model.Throw(c, errhandle.UserNonExists)
		return
	}
	model.OK[any](c, nil)
}

/* ----------删除用户-------- */
// 删除指定管理员 godoc
// @Summary 删除指定管理员
// @Schemes
// @Description 删除指定管理员
// @Tags admin
// @Accept json
// @Produce json
// @Param   Authorization     header    string  true   "登录返回的Token"
// @Param   DeleteUser     body     model.UserIDOnlyRequest  true  "需要删除的管理员ID"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/deleteAdmin [delete]
func (t *Admin) DeleteAdmin(c *gin.Context) {
	info, ok := c.Get("info")
	if !ok {
		model.ThrowWithMsg(c, "获取上下文中的用户信息失败")
		return
	}
	userInfo := info.(*model.UserInfo)
	b, err := c.GetRawData()
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	var req model.UserIDOnlyRequest
	json.Unmarshal(b, &req)

	var u model.User
	err = global.GVB_DB.Transaction(func(tx *gorm.DB) error {
		//TODO: 还要删除关联的文章评论等
		result := tx.Table("users").
			Where("id = ? AND role = ?", req.UserID, "admin").
			Delete(&u)
		if result.Error != nil {
			return errors.New(errhandle.UserNonExists.String())
		}
		if result.RowsAffected == 0 {
			// 没有任何记录被更新，可能是因为用户不存在或角色不匹配
			return errors.New(errhandle.UserNonExists.String())
		}
		return nil
	})
	if err != nil {
		global.GVB_LOGGER.Error(err)
		model.ThrowError(c, err)
		return
	}
	// 如果要删除的管理员刚好是当前在线的
	if req.UserID == userInfo.UserID {
		global.GVB_REDIS.Del(cb, userInfo.UUID)
	}
	model.OK[any](c, nil)
}

// 删除指定用户 godoc
// @Summary 删除指定用户
// @Schemes
// @Description 删除指定用户
// @Tags admin
// @Accept json
// @Produce json
// @Param   Authorization     header    string  true   "登录返回的Token"
// @Param   DeleteUser     body     model.UserIDOnlyRequest  true  "需要删除的用户ID"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/deleteUser [delete]
func (t *Admin) DeleteUser(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	var req model.UserIDOnlyRequest
	json.Unmarshal(b, &req)

	var u model.User
	err = global.GVB_DB.Transaction(func(tx *gorm.DB) error {
		//TODO: 还要删除关联的文章评论等
		result := tx.Table("users").
			Where("id = ?", req.UserID).
			Delete(&u)
		if result.Error != nil {
			return errors.New(errhandle.UserNonExists.String())
		}
		if result.RowsAffected == 0 {
			// 没有任何记录被更新，可能是因为用户不存在或角色不匹配
			return errors.New(errhandle.UserNonExists.String())
		}
		return nil
	})
	if err != nil {
		global.GVB_LOGGER.Error(err)
		model.ThrowError(c, err)
		return
	}
	model.OK[any](c, nil)
}
