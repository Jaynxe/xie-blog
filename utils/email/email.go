package email

import (
	"github.com/Jaynxe/xie-blog/global"
	"gopkg.in/gomail.v2"
)

// EmailSubject 表示邮件的主题类型
type EmailSubject string

// 定义不同类型的邮件主题
const (
	VerificationCode EmailSubject = "平台验证码"
	OperationNotice  EmailSubject = "操作通知"
	AlarmNotice      EmailSubject = "告警通知"
)

// EmailApi 结构体包含邮件主题
type EmailApi struct {
	Subject EmailSubject
}

// Send 发送邮件
// recipient: 收件人邮箱地址
// body: 邮件内容
func (api EmailApi) Send(recipient, body string) error {
	return sendEmail(recipient, string(api.Subject), body)
}

// NewVerificationCodeApi 创建验证码邮件的API实例
func NewVerificationCodeApi() EmailApi {
	return EmailApi{
		Subject: VerificationCode,
	}
}

// NewOperationNoticeApi 创建操作通知邮件的API实例
func NewOperationNoticeApi() EmailApi {
	return EmailApi{
		Subject: OperationNotice,
	}
}

// NewAlarmNoticeApi 创建告警通知邮件的API实例
func NewAlarmNoticeApi() EmailApi {
	return EmailApi{
		Subject: AlarmNotice,
	}
}

// sendEmail 发送邮件的具体实现
// recipient: 收件人邮箱地址
// subject: 邮件主题
// body: 邮件内容
func sendEmail(recipient, subject, body string) error {
	emailConfig := global.GVB_CONFIG.Email
	return sendMail(
		emailConfig.SenderEmail,
		emailConfig.Password,
		emailConfig.Host,
		emailConfig.Port,
		recipient,
		emailConfig.DefaultFromEmail,
		subject,
		body,
	)
}

// sendMail 使用gomail发送邮件
// senderEmail: 发件人邮箱地址
// authCode: 发件人邮箱的授权码
// host: SMTP服务器地址
// port: SMTP服务器端口
// recipient: 收件人邮箱地址
// senderName: 发件人显示名称
// subject: 邮件主题
// body: 邮件内容
func sendMail(senderEmail, authCode, host string, port int, recipient, senderName, subject, body string) error {
	// 创建新的邮件消息
	message := gomail.NewMessage()
	message.SetHeader("From", message.FormatAddress(senderEmail, senderName)) // 设置发件人
	message.SetHeader("To", recipient)                                        // 设置收件人
	message.SetHeader("Subject", subject)                                     // 设置邮件主题
	message.SetBody("text/html", body)                                        // 设置邮件正文

	// 创建SMTP拨号器
	dialer := gomail.NewDialer(host, port, senderEmail, authCode)
	// 发送邮件
	err := dialer.DialAndSend(message)
	return err
}
