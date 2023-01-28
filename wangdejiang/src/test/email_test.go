package test

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "aei <imaei@foxmail.com>"
	e.To = []string{"449857252@qq.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("<h1>你好！</h1></br><p>您的验证码是<b>123456</b></p>")
	// 返回EOF需要关闭SSL
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("",
		"3035816700@qq.com", "sdjhqnbaxexuddgd", "smtp.qq.com"),
		&tls.Config{
			ServerName:         "smtp.qq.com",
			InsecureSkipVerify: true,
		},
	)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
