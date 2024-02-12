package email

import (
	"fmt"
	. "github.com/burak-aydin/website/config"
	log "github.com/sirupsen/logrus"
	"net/smtp"
)

func SendEmail(content string) {

	auth := smtp.PlainAuth("", GetEnvironmentVariable(EnvKeyEmailSender), GetEnvironmentVariable(EnvKeyEmailPassword), GetEnvironmentVariable(EnvKeyEmailHost))
	// Here we do it all: connect to our server, set up a message and send it

	msg := []byte(fmt.Sprintf("From: App %s\r\n", GetEnvironmentVariable(EnvKeyEmailSender)) +
		fmt.Sprintf("To: Me %s\r\n", GetEnvironmentVariable(EnvKeyEmailReceiver)) +
		"Subject: You have received a message\r\n" +
		"\r\n" +
		content +
		"\r\n")

	err := smtp.SendMail(fmt.Sprintf("%s:587", GetEnvironmentVariable(EnvKeyEmailHost)), auth, GetEnvironmentVariable(EnvKeyEmailSender), []string{GetEnvironmentVariable(EnvKeyEmailReceiver)}, msg)
	if err != nil {
		log.WithError(err).Error("Error while sending email")
	}
	log.Info("Email Sent Successfully!")
}
