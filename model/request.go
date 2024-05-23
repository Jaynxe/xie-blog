package model

type QueryRequest struct {
	UserID   int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	NickName string `json:"nick_name,omitempty"`
}
type UserLoginRequest struct {
	Name     string `json:"name,omitempty" binding:"required" msg:"用户名"`
	Password string `json:"password,omitempty" binding:"required" msg:"密码"`
}
type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	NickName string `json:"nick_name"`
	Sex      string `json:"sex"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

// 图片
// 分页请求
type PageRequest struct {
	Page  int    `form:"page" json:"page,omitempty"`   //页码
	Key   string `form:"key" json:"key,omitempty"`     //关键字
	Limit int    `form:"limit" json:"limit,omitempty"` //每页大小
	Sort  string `form:"sort" json:"sort,omitempty"`   //排序(升序/降序)
}
type DelIdListRequest struct {
	IdList []int64 `json:"id_list" binding:"required" msg:"请输入要删除的内容id列表"`
}
type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择图片id"` //必传
	Name string `json:"name" binding:"required" msg:"请输入图片名称"`
}

// 菜单
type MenuRequest struct {
	Title    string `json:"title,omitempty" binding:"required" msg:"菜单标题"`
	URL      string `json:"url,omitempty" binding:"required" msg:"菜单url"`
	ParentID *int   `json:"parent_id,omitempty"`
	Sort     int    `json:"sort,omitempty" binding:"required" msg:"菜单的排序"`
}
type MenuUptateRequest struct {
	Title    string `json:"title,omitempty" binding:"required" msg:"菜单标题"`
	URL      string `json:"url,omitempty" binding:"required" msg:"菜单url"`
	ParentID *int   `json:"parent_id,omitempty"`
	Sort     int    `json:"sort,omitempty" binding:"required" msg:"菜单的排序"`
}

// 用户
type UserIDOnlyRequest struct {
	UserID int64 `json:"id"`
}
type ModifyUserRequest struct {
	UserIDOnlyRequest
	Name     string `json:"name,omitempty"`
	NickName string `json:"nick_name,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Email    string `json:"email,omitempty"`
}

type ModifyPasswordRequest struct {
	UserIDOnlyRequest
	OldPwd string `json:"old_pwd,omitempty"`
	NewPwb string `json:"new_pwb,omitempty"`
}

type LoginWithEmailRequest struct {
	Email    string  `json:"email" binding:"required" msg:"邮箱非法"`
	Password string  `json:"password"`
	Code     *string `json:"code"`
}
type TagRequest struct {
	Name string `json:"Name,omitempty" binding:"required" msg:"请输入标签名"`
}
