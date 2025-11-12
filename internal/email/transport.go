package email

import (
	"context"
	"time"

	"github.com/carfdev/carfdev-emailsvc/internal/config"
	"github.com/carfdev/carfdev-emailsvc/internal/logger"
	"github.com/carfdev/carfdev-emailsvc/internal/natsx"
	"github.com/carfdev/carfdev-emailsvc/internal/types"
	"github.com/carfdev/carfdev-emailsvc/internal/util"
	"github.com/nats-io/nats.go"
)

const (
	SubjectSendContact = "email.send_contact"
)

const (
	CodeBadRequest = "bad_request"
	CodeNotFound   = "not_found"
	CodeInvalid    = "invalid"
	CodeInternal   = "internal"
)

type Transport struct {
	bus     *natsx.Bus
	service Service
	lg      *logger.Logger
}

func NewTransport(bus *natsx.Bus, svc Service, lg *logger.Logger) *Transport {
	return &Transport{bus: bus, service: svc, lg: lg}
}

type Envelope struct {
	Data  any  `json:"data,omitempty"`
	Error *Err `json:"error,omitempty"`
}

type Err struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ok(v any) *Envelope              { return &Envelope{Data: v} }
func fail(code, msg string) *Envelope { return &Envelope{Error: &Err{Code: code, Message: msg}} }

func (t *Transport) RegisterHandlers(cfg *config.Config) error {
	subjects := map[string]nats.MsgHandler{
		SubjectSendContact: t.handleSendContact(),
	}

	for subj, handler := range subjects {
		if _, err := t.bus.QueueSubscribe(subj, cfg.Service, handler); err != nil {
			return err
		}
		t.lg.Infow("registered NATS handler", "subject", subj)
	}

	return nil
}

func (t *Transport) handleSendContact() nats.MsgHandler {
	return func(m *nats.Msg) {
		var in *types.SendContactRequest
		if err := util.StrictUnmarshal(m.Data, &in); err != nil {
			t.lg.Errorw("invalid payload handleSendContact", "error", err)
			t.bus.Reply(m, fail(CodeBadRequest, err.Error()))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		a, err := t.service.SendContact(ctx, in)
		if err != nil {
			t.lg.Errorw("failed to send contact email", "error", err)
			t.bus.Reply(m, fail(CodeInternal, err.Error()))
			return
		}

		t.bus.Reply(m, ok(a))
	}
}
