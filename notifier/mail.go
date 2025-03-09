package notifier

import (
	"fmt"
	"net/smtp"
)
// EmailNotifier is a struct that holds the configuration for sending emails.
type EmailNotifier struct {
	SMTPServer string
	Port       string
	Username   string
	Password   string
	To         []string
}
// Notify sends an email with the given message.
func (e *EmailNotifier) Notify(message string) {
	auth := smtp.PlainAuth("", e.Username, e.Password, e.SMTPServer)
	to := e.To
	subject := "Log Notification"
	body := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, message)

	err := smtp.SendMail(e.SMTPServer+":"+e.Port, auth, e.Username, to, []byte(body))
	if err != nil {
		fmt.Printf("Failed to send email: %s\n", err)
	}
}