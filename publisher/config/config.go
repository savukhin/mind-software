package config

type Config struct {
	NatsUrl string `envconfig:"NATS_URL" required:"true"`
}
