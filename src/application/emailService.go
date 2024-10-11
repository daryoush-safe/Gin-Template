package application

import (
	"first-project/src/bootstrap"
	"fmt"
	"net/smtp"
)

type EmailService struct {
	emailInfo *bootstrap.EmailInfo
}

func NewEmailService(emailInfo *bootstrap.EmailInfo) *EmailService {
	return &EmailService{
		emailInfo: emailInfo,
	}
}

// TODO: change this to communication and then email sender!
func (emailService *EmailService) SendVerificationEmail(username string, toEmail string, token string) {
	from := emailService.emailInfo.EmailFrom
	password := emailService.emailInfo.EmailPassword
	smtpHost := emailService.emailInfo.SMTPHost
	smtpPort := emailService.emailInfo.SMTPPort

	// TODO: remove jwt use otp instead
	verificationLink := (fmt.Sprintf("http://localhost:8080/v1/verifyEmail/%s", token))
	// TODO: translation
	message := []byte(
		fmt.Sprintf(
			"dear %s;\nPlease verify your email by clicking on the link:\n%s",
			username, verificationLink),
	)
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, message)
	if err != nil {
		panic(err)
	}
}
