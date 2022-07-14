package tm

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"net/smtp"
	"net/textproto"
)

func sendMail(subject string, to string, html string) error {
	var from = viper.GetString("MAIL_FORM")
	e := &email.Email{
		To:      []string{to},
		From:    fmt.Sprintf("%s <%s>", viper.GetString("MAIL_NAME"), from),
		Subject: subject,
		HTML:    []byte(html),
		Headers: textproto.MIMEHeader{},
	}
	var host = viper.GetString("MAIL_HOST")
	var port = viper.GetString("MAIL_PORT")
	var password = viper.GetString("MAIL_USERNAME")
	var username = viper.GetString("MAIL_PASSWORD")
	var addr = fmt.Sprintf("%s:%s", host, port)
	err := e.Send(addr, smtp.PlainAuth("", username, password, host))
	if err != nil {
		return err
	}
	return nil
}
