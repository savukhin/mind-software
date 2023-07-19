package config

const ReleaseMode = "release"
const DebugMode = "debug"

type Config struct {
	PostgresIPs         []string `envconfig:"POSTGRES_IPS" required:"true"`
	PostgresPort        int      `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresUser        string   `envconfig:"POSTGRES_USER" default:"mindsoftware"`
	PostgresPassword    string   `envconfig:"POSTGRES_PASSWORD" default:"mindsoftware"`
	PostgresDB          string   `envconfig:"POSTGRES_DB" default:"mindsoftware"`
	PostgresAutoMigrate bool     `envconfig:"POSTGRES_AUTOMIGRATE" default:"true"`

	NatsUrl           string `envconfig:"NATS_URL" required:"true"`
	NatsLogsTopic     string `envconfig:"NATS_LOGS_TOPIC" default:"logs"`
	NatsQueryTopic    string `envconfig:"NATS_QUERY_TOPIC" default:"query"`
	NatsResponseTopic string `envconfig:"NATS_RESPONSE_TOPIC" default:"response"`

	Mode string `envconfig:"MODE" default:"realease"`
}
