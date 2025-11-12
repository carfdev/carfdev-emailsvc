package natsx

import (
	"encoding/json"
	"time"

	"github.com/carfdev/carfdev-emailsvc/internal/logger"
	"github.com/nats-io/nats.go"
)

type Bus struct {
	nc *nats.Conn
	lg *logger.Logger
}

func Connect(url string, name string, lg *logger.Logger) (*Bus, error) {
	opts := []nats.Option{
		nats.Name(name),
		nats.ReconnectWait(2 * time.Second),
		nats.MaxReconnects(-1),
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			lg.Errorw("NATS disconnected", "error", err)
		}),
		nats.ReconnectHandler(func(_ *nats.Conn) {
			lg.Infow("NATS reconnected")
		}),
		nats.ClosedHandler(func(_ *nats.Conn) {
			lg.Errorw("NATS connection closed")
		}),
	}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, err
	}

	lg.Infow("connected to NATS", "url", url)
	return &Bus{nc: nc, lg: lg}, nil
}

func (b *Bus) QueueSubscribe(subject, queue string, fn nats.MsgHandler) (*nats.Subscription, error) {
	return b.nc.QueueSubscribe(subject, queue, fn)
}

func (b *Bus) Reply(m *nats.Msg, v any) {
	data, err := json.Marshal(v)
	if err != nil {
		b.lg.Errorw("failed to marshal NATS reply", "error", err)
		return
	}
	if err := m.Respond(data); err != nil {
		b.lg.Errorw("failed to send NATS reply", "error", err)
	}
}

func (b *Bus) Publish(subject string, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return b.nc.Publish(subject, data)
}

func (b *Bus) Close() {
	if b.nc != nil && !b.nc.IsClosed() {
		b.nc.Close()
		b.lg.Infow("NATS connection closed")
	}
}
