package service

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

// SendCode 发送验证码 toEmail 邮箱
func SendCode(toEmail, code string) (err error) {
	e := email.NewEmail()
	e.From = "aei <imaei@foxmail.com>"
	e.To = []string{toEmail}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("<h1>你好！</h1></br><p>您的验证码是<b>" + code + "</b></p>")
	// 返回EOF需要关闭SSL
	err = e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("",
		"3035816700@qq.com", "sdjhqnbaxexuddgd", "smtp.qq.com"),
		&tls.Config{
			ServerName:         "smtp.qq.com",
			InsecureSkipVerify: true,
		},
	)
	return
}

func GenarateCode() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}
