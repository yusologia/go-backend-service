package config

import (
	"github.com/yusologia/go-core/v2"
	"os"
)

var (
	SMTPMail logiacore.MailConf
)

func InitMail() {
	SMTPMail = logiacore.MailConf{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     os.Getenv("MAIL_PORT"),
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}
}
