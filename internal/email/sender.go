package email

import (
	"fmt"

	"github.com/carfdev/carfdev-emailsvc/internal/config"
	"github.com/resend/resend-go/v3"
)

func Sender(to []string, subject string, message string, cfg *config.Config) (*resend.SendEmailResponse, error) {
	client := resend.NewClient(cfg.Key)

	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("Carfdev <%s>", cfg.From),
		To:      to,
		Html:    message,
		Subject: subject,
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		return nil, err
	}
	return sent, nil
}
