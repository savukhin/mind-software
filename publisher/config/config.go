package config

type Config struct {
	NatsUrl       string `envconfig:"NATS_URL" required:"true"`
	QueriesMin    int    `envconfig:"QUERIES_MIN" default:"1"`
	NatsLogsTopic string `envconfig:"NATS_LOGS_TOPIC" default:"logs"`
}
