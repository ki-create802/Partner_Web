package services

//定义发送验证码到用户电子邮件的服务
import (
	"Partner_Web/Partner_Server/config" // 引入config包
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendVerificationCode(email string, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.SMTPUsername) // 使用config.go中的SMTP用户名作为发件人
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset Verification Code")
	m.SetBody("text/plain", fmt.Sprintf("Your password reset verification code is: %s", code))

	d := gomail.NewDialer(
		config.SMTPServer, // 使用config.go中的SMTP服务器地址
		config.SMTPPort,   // 使用config.go中的SMTP端口
		config.SMTPUsername,
		config.SMTPPassword,
	)
	return d.DialAndSend(m)
}
