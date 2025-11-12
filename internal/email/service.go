package email

import (
	"context"
	"errors"
	"fmt"

	"github.com/carfdev/carfdev-emailsvc/internal/config"
	"github.com/carfdev/carfdev-emailsvc/internal/logger"
	"github.com/carfdev/carfdev-emailsvc/internal/template"
	"github.com/carfdev/carfdev-emailsvc/internal/types"
)

type Service interface {
	SendContact(ctx context.Context, req *types.SendContactRequest) (*SendResponse, error)
}

type service struct {
	lg  *logger.Logger
	cfg *config.Config
}

var (
	ErrInvalidPayload = errors.New("invalid payload")
)

func NewService(lg *logger.Logger, cfg *config.Config) Service {
	return &service{lg: lg, cfg: cfg}
}

func (s *service) SendContact(ctx context.Context, req *types.SendContactRequest) (*SendResponse, error) {

	message := template.ContactRequestTemplate(req)

	sent, err := Sender([]string{s.cfg.Admin}, "New Contact Request", message, s.cfg)
	if err != nil {
		return nil, err
	}

	return &SendResponse{
		Status:  200,
		Message: fmt.Sprintf("Email sent, ID: %s", sent.Id),
	}, nil
}
