package tm

import (
	"errors"
	"fmt"
	comm2 "github.com/glide-im/api/internal/api/comm"
	"github.com/glide-im/api/internal/pkg/db"
	email2 "github.com/jordan-wright/email"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"html/template"
	"net/smtp"
	"net/textproto"
	"regexp"
	"time"
)

var (
	VerifyCodeError = comm2.NewApiBizError(2001, "验证码错误请重试")
	VerifyCodeLimit = errors.New("获取验证码频率太高了，请稍后重试")
)

var (
	Limit    = 3 // 一个小时最多3个
	Expire   = 30 * time.Minute
	CacheKey = "user:verify:"
)

var VerifyCodeU = &VerifyCode{}

type VerifyCode struct{}

/** 发送验证码 */
func (c *VerifyCode) SendVerifyCode(value string, tm string) (err error) {
	if c.isLimit(value) {
		return VerifyCodeLimit
	}
	var key = c.getKey(value)
	var verifyCode = cast.ToString(RandomInt(4))

	matched, err := regexp.Match("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$", []byte(value))
	if err != nil {
		fmt.Println(err)
		return err
	}

	if matched {
		err = c.SendEmail(tm, value, verifyCode)
	} else {
		err = c.sendPhone(tm, value, verifyCode)
	}
	if err != nil {
		return err
	}

	err = c.setVerifyCode(value, verifyCode)
	defer db.Redis.Expire(key, Expire)
	if err != nil {
		return err
	}
	return nil
}

func (c *VerifyCode) ValidateVerifyCode(value string, code string) error {
	rs, err := db.Redis.HGet(c.getKey(value), code).Result()
	if err != nil {
		return VerifyCodeError
	}
	if len(rs) == 0 {
		return VerifyCodeError
	}
	return nil
}

func (c *VerifyCode) ClearLimit(value string) {
	db.Redis.Del(c.getKey(value))
}

func (c *VerifyCode) setVerifyCode(value string, code string) error {
	_, err := db.Redis.HSet(c.getKey(value), code, 1).Result()
	return err
}

func (c *VerifyCode) isLimit(value string) bool {
	rs, err := db.Redis.HGetAll(c.getKey(value)).Result()
	if err != nil {
		return false
	}
	return len(rs) >= Limit
}

func (c *VerifyCode) getKey(value string) string {
	return CacheKey + value
}

func (c *VerifyCode) SendEmail(tm string, email string, code string) error {
	res, err := template.ParseFiles(tm)
	if err != nil {
		fmt.Println(err)
	}
	captcha := Captcha{
		Code: code,
	}
	var html Writer
	_ = res.Execute(&html, captcha)

	var from = viper.GetString("Mail.MAIL_FORM")
	fmt.Println(html.Html)
	e := &email2.Email{
		To:      []string{email},
		From:    fmt.Sprintf("%s <%s>", viper.GetString("Mail.MAIL_NAME"), from),
		Subject: "验证码登录",
		HTML:    []byte(html.Html),
		Headers: textproto.MIMEHeader{},
	}
	var host = viper.GetString("Mail.MAIL_HOST")
	var port = viper.GetString("Mail.MAIL_PORT")
	var password = viper.GetString("Mail.MAIL_PASSWORD")
	var username = viper.GetString("Mail.MAIL_USERNAME")
	var addr = fmt.Sprintf("%s:%s", host, port)
	fmt.Println(addr, username, password, host)
	err = e.Send(addr, smtp.PlainAuth("ssl", username, password, host))
	if err != nil {
		return err
	}
	return nil

}

func (c *VerifyCode) sendPhone(tm string, email string, code string) error {
	return nil
}

func (c *VerifyCode) loadTemplateHtml(template string) {

}
