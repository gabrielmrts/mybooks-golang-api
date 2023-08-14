package services

import (
	"bytes"
	"fmt"
	"os"

	"text/template"

	"github.com/gabrielmrts/mybooks-golang-api/config"
	"github.com/gabrielmrts/mybooks-golang-api/internal/utils"
	"gopkg.in/gomail.v2"
)

type MailService struct{}

func NewMailService() *MailService {
	return &MailService{}
}

func (ms MailService) SendMail(subject string, to string, templateName string, variables interface{}) error {
	config := config.GetConfig()
	smtpHost := config.SMTP_HOST
	smtpPort := config.SMTP_PORT

	// Settings for auth in the mail server
	dialer := gomail.NewDialer(smtpHost, smtpPort, "", "")

	message := gomail.NewMessage()
	message.SetHeader("From", "mybooks@support.com")
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)

	templatesDir := utils.GetTemplatesFolder()
	templateContent, err := os.ReadFile(fmt.Sprintf("%s/%s.html", templatesDir, templateName))

	if err != nil {
		return err
	}

	tmpl, err := template.New("emailTemplate").Parse(string(templateContent))

	if err != nil {
		return err
	}

	var emailBody bytes.Buffer
	if err := tmpl.Execute(&emailBody, variables); err != nil {
		return err
	}

	message.SetBody("text/html", emailBody.String())

	if err := dialer.DialAndSend(message); err != nil {
		return err
	} else {
		return nil
	}
}
