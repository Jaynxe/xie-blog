package flag

import (
	"fmt"
	"net/mail"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/Jaynxe/xie-blog/utils"
	"github.com/Jaynxe/xie-blog/utils/pwd"
	"github.com/Jaynxe/xie-blog/utils/snowflake"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// CreateUser 命令行创建用户
func CreateUser(role string) {
	//用户名，昵称，密码，确认密码，邮箱
	var (
		userName   string
		nickName   string
		passWord   string
		rePassWord string
		email      string
		sex        string
	)
	var user model.User
UN:
	fmt.Printf("请输入用户名:")
	fmt.Scan(&userName)
	//判断用户名是否存在
	err := global.GVB_DB.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)}).Take(&user, "user_name = ?", userName).Error
	if err == nil {
		global.GVB_LOGGER.Warn("用户名已经存在，请重新输入")
		goto UN
	}
	fmt.Printf("请输入昵称:")
	fmt.Scan(&nickName)
	fmt.Printf("请输入性别:")
	fmt.Scan(&sex)
EM:
	fmt.Printf("请输入邮箱:")
	fmt.Scan(&email)
	_, err = mail.ParseAddress(email)
	if err != nil {
		global.GVB_LOGGER.Warn("邮箱格式不正确，请重新输入")
		goto EM
	}
PW:
	fmt.Printf("请输入密码:")
	fmt.Scan(&passWord)
	isValid := utils.IsValidPassword(passWord)
	if !isValid {
		global.GVB_LOGGER.Warn("密码长度要大等于8且包含大小写字母, 请重新输入")
		goto PW
	}
	fmt.Printf("请输入确认密码:")
	fmt.Scan(&rePassWord)
	if passWord != rePassWord {
		global.GVB_LOGGER.Warn("两次密码不一致, 请重新输入")
		goto PW
	}

	//对密码进行hash
	salt := pwd.HashAndSalt(passWord)

	err = global.GVB_DB.Create(&model.User{
		ID:       snowflake.ID(),
		NickName: nickName,
		Name:     userName,
		Password: salt,
		Email:    email,
		Role:     role,
		Sex:      sex,
	}).Error
	if err != nil {
		global.GVB_LOGGER.Error(err)
		return
	}
	if role == "admin" {
		global.GVB_LOGGER.Infof("管理员%s创建成功", userName)
	} else {
		global.GVB_LOGGER.Infof("用户%s创建成功", userName)
	}
}
