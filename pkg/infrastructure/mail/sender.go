package mail

import (
	"context"
	"exchange_rate/pkg/domain"
	"fmt"
	"net/smtp"
	"os"
)

type EmailSender struct {
	config  *Config
	auth    smtp.Auth
	address string
	subject string
}

func NewEmailService() (*EmailSender, error) {
	mailConfig, err := newConfig(
		os.Getenv("EMAIL_ADDRESS"),
		os.Getenv("EMAIL_APP_CODE"),
		os.Getenv("EMAIL_SUBJECT"),
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
	)
	if err != nil {
		return nil, err
	}

	return &EmailSender{
		config:  mailConfig,
		auth:    smtp.PlainAuth("", mailConfig.address, mailConfig.appKey, mailConfig.smtpHost),
		address: fmt.Sprintf("%s:%s", mailConfig.smtpHost, mailConfig.smtpPort),
	}, nil
}

func (e *EmailSender) SendEmail(
	_ context.Context, data *domain.CurrencyRate, receivers ...string,
) error {
	return e.sendEmail(e.crateTemplate(data), receivers...)
}

func (e *EmailSender) sendEmail(message []byte, receiversEmail ...string) error {
	return smtp.SendMail(
		e.address,
		e.auth,
		e.config.address,
		receiversEmail,
		message,
	)
}
