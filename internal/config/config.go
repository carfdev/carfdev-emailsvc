package config

import (
	"log"
	"os"
)

type Config struct {
	NatsURL string
	Service string
	From    string
	Env     string
	Key     string
	Admin   string
}

func Load() *Config {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = "nats://localhost:4222"
	}

	service := os.Getenv("SERVICE_NAME")
	if service == "" {
		service = "emailsvc"
	}

	from := os.Getenv("EMAIL_FROM")
	if from == "" {
		from = "noreply@example.com"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	key := os.Getenv("EMAIL_KEY")
	if key == "" {
		log.Fatal("⚠️  EMAIL_KEY is not configured. Please set the environment variable.")
	}

	admin := os.Getenv("EMAIL_ADMIN")
	if admin == "" {
		admin = "admin@example.com"
	}

	return &Config{
		NatsURL: natsURL,
		Service: service,
		From:    from,
		Env:     env,
		Key:     key,
		Admin:   admin,
	}
}
