package config

type Config struct {
	NatsUrl       string `envconfig:"NATS_URL" required:"true"`
	NatsLogsTopic string `envconfig:"NATS_LOGS_TOPIC" default:"logs"`
	QueriesMin    int    `envconfig:"QUERIES_MIN" default:"1"`
}
