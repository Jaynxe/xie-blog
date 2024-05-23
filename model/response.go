package model

import (
	"net/http"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/utils/errhandle"
	"github.com/Jaynxe/xie-blog/utils/valid"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	UserID int64  `json:"userid"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Role   string `json:"role"`
}

type CommonResponse[T any] struct {
	Data   T                 `json:"data,omitempty"`
	Status errhandle.ErrCode `json:"status"`
	Reason string            `json:"reason,omitempty"`
	Msg    string            `json:"msg,omitempty"`
}
type TokenResponse struct {
	AccessToken  string `json:"token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	ExpiredAt    int64  `json:"expiredAt,omitempty"`
}

type GetUserResponse struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	NickName string `json:"nick_name,omitempty"`
	Sex      string `json:"sex,omitempty"`
}
type ImageResponse struct {
	FilePath     string `json:"file_path,omitempty"`
	IsSucceed    bool   `json:"is_succeed,omitempty"`
	UploadStatus string `json:"upload_status,omitempty"`
}

// 服务器内部错误
func ThrowError(c *gin.Context, err error) {
	global.GVB_LOGGER.Errorln(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, CommonResponse[any]{
		Status: errhandle.InnerError,
		Reason: err.Error(),
	})
}

// 已知的定义错误
func Throw(c *gin.Context, errCode errhandle.ErrCode) {
	c.AbortWithStatusJSON(http.StatusBadRequest, CommonResponse[any]{
		Status: errCode,
		Reason: errCode.String(),
	})
}

// 返回自定义的错误字符串
func ThrowWithMsg(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, CommonResponse[any]{
		Status: errhandle.OtherError,
		Msg:    msg,
	})
}

// 数据绑定错误
func ThrowBindError(c *gin.Context, obj any, err error) {
	msg := valid.GetValidMsg(err, obj)
	ThrowWithMsg(c, msg)
}

// 成功
func OK[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, CommonResponse[T]{
		Data: data,
	})
}

// 成功并携带给定的信息
func OKWithMsg[T any](c *gin.Context, data T, msg string) {
	c.JSON(http.StatusOK, CommonResponse[T]{
		Data: data,
		Msg:  msg,
	})
}
