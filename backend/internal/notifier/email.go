package notifier

import (
	"log"
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	from := "your-email@example.com"
	password := "your-email-password"
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	message := []byte("Subject: " + subject + "\r\n\r\n" + body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}