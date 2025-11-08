package utils

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to string, subject string, body string) error {
	from := "bookshop430@gmail.com"
	appPassword := "iewaxjyygwxxyzba" // No spaces here!

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		from, to, subject, body)

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	err := smtp.SendMail(smtpHost+":"+smtpPort,
		smtp.PlainAuth("", from, appPassword, smtpHost),
		from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}
