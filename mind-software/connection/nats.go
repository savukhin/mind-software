package connection

import (
	"mind-software/config"

	"github.com/nats-io/nats.go"
)

func ConnectNats(cfg *config.Config) (*nats.Conn, error) {
	return nats.Connect(cfg.NatsUrl)
}
