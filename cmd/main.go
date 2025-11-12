package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/carfdev/carfdev-emailsvc/internal/config"
	"github.com/carfdev/carfdev-emailsvc/internal/email"
	"github.com/carfdev/carfdev-emailsvc/internal/logger"
	"github.com/carfdev/carfdev-emailsvc/internal/natsx"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg := config.Load()
	lg := logger.New(cfg.Env)
	lg.Infow("starting service", "service", cfg.Service, "env", cfg.Env)

	// NATS
	bus, err := natsx.Connect(cfg.NatsURL, cfg.Service, lg)
	if err != nil {
		lg.Fatalw("nats connection failed", "error", err)
	}
	defer bus.Close()

	service := email.NewService(lg, cfg)
	lg.Infow("email service initialized")
	transport := email.NewTransport(bus, service, lg)
	lg.Infow("transport initialized")
	if err := transport.RegisterHandlers(cfg); err != nil {
		lg.Fatalw("register handlers failed", "error", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	<-ctx.Done()
	stop()
	lg.Infow("shutting down")
	bus.Close()
	lg.Infow("NATS connection closed")
}
