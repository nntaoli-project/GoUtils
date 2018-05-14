package email

import "testing"

func TestSendEmail(t *testing.T) {
	SendEmail(Sender{"test@163.com", "123456", "smtp.163.com", 25},
		"xxx@163.com", "测试发送邮件", "测试发邮件", false)
}

func TestSendEmailWithTls(t *testing.T) {
	SendEmailWithTls(Sender{"test@163.com", "123456", "smtp.163.com", 465},
		"xxx@163.com", "测试发邮件", "<b>测试发邮件</b>", true)
}
