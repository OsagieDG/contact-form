package handlers

import (
	"fmt"
	"net/smtp"
	"os"
)

func sendEmail(name, email, phoneNumber, message string) error {
	from := os.Getenv("SMTP_FROM_EMAIL")
	to := "dgodstand@gmail.com"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	smtpPassword := os.Getenv("SMTP_PASSWORD") // generate your personal SMTP password using Google 2-factor authentication.
	auth := smtp.PlainAuth("", from, smtpPassword, smtpHost)

	subject := "CONTACT_REQUEST"
	body := fmt.Sprintf("Name: %s\r\nEmail: %s\r\nPhone Number: %s\r\nMessage: %s", name, email, phoneNumber, message)

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
}
