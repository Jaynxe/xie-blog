package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/Jaynxe/xie-blog/utils"
	"github.com/Jaynxe/xie-blog/utils/errhandle"
	"github.com/Jaynxe/xie-blog/utils/pwd"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	cb = context.Background()
)

func (u *User) UserInfo(c *gin.Context) {

}

// 用户注销 godoc
// @Summary 用户注销
// @Schemes
// @Description 用户注销
// @Tags user
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[string]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/user/logout [post]
func (u *User) Logout(c *gin.Context) {
	info, ok := c.Get("info")
	if !ok {
		model.ThrowWithMsg(c, "获取token中的用户信息失败")
		return
	}
	userInfo := info.(*model.UserInfo)
	_, err := global.GVB_REDIS.Del(cb, userInfo.UUID).Result()
	if err != nil {
		model.ThrowWithMsg(c, "用户注销失败")
		return
	}
	model.OK(c, "用户注销成功")

}

// 用户删除 godoc
// @Summary 用户删除
// @Schemes
// @Description 用户删除
// @Tags user
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[string]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/user/deleteUser [delete]
func (u *User) DeleteUser(c *gin.Context) {
	info, ok := c.Get("info")
	if !ok {
		model.ThrowWithMsg(c, "获取token中的用户信息失败")
		return
	}
	userInfo := info.(*model.UserInfo)
	var user model.User
	err := global.GVB_DB.Transaction(func(tx *gorm.DB) error {
		// TODO:关联的文章评论也要删除
		err := tx.Delete(&user, "id = ?", userInfo.UserID).Error
		if err != nil {
			global.GVB_LOGGER.Error(err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		global.GVB_LOGGER.Error(err.Error())
		model.ThrowError(c, err)
		return
	}

	global.GVB_REDIS.Del(cb, userInfo.UUID)

	model.OK(c, "删除用户成功")
}

// 修改用户密码 godoc
// @Summary 修改用户密码
// @Schemes
// @Description 修改用户密码
// @Tags user
// @Accept json
// @Produce json
// @Param   Authorization     header    string  true   "登录返回的Token"
// @Param   ModifyPasswordRequest    body   model.ModifyPasswordRequest  true  "id,新密码"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/user/modifyUserPassword [patch]
func (u *User) ModifyUserPassword(c *gin.Context) {
	info, ok := c.Get("info")
	if !ok {
		model.ThrowWithMsg(c, "获取token中的用户信息失败")
		return
	}
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

	userInfo := info.(*model.UserInfo)
	global.GVB_REDIS.Del(cb, userInfo.UUID)

	model.OK(c, "密码修改成功")
}

// 修改用户信息 godoc
// @Summary 修改用户信息
// @Schemes
// @Description 修改用户信息
// @Tags user
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   ModifyUserRequest     body     model.ModifyUserRequest  true  "要修改的用户的信息"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/user/modifyUser [patch]
func (u *User) ModifyUser(c *gin.Context) {
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

// 获取当前用户信息 godoc
// @Summary 获取当前用户信息
// @Schemes
// @Description 获取当前用户信息
// @Tags user
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[string]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/user/getUserInfo [get]
func (u *User) GetUserInfo(c *gin.Context) {
	info, ok := c.Get("info")
	if !ok {
		model.ThrowWithMsg(c, "获取token中的用户信息失败")
		return
	}
	userInfo := info.(*model.UserInfo)
	var user model.User
	err := global.GVB_DB.Omit("password").First(&user, userInfo.UserID).Error
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	model.OKWithMsg(c, user, fmt.Sprintf("获取用户[%s]信息成功", userInfo.Name))
}
