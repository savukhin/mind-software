package connection

import (
	"fmt"
	"mind-software/config"
	"mind-software/pgmodels"

	"github.com/serialx/hashring"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresClient struct {
	client *gorm.DB
	host   string
}

func NewPostgresClient(client *gorm.DB, host string) *postgresClient {
	return &postgresClient{
		client: client,
		host:   host,
	}
}

type IPostgresClients interface {
	GetShard(key string) *gorm.DB
}

type postgresClients struct {
	clients map[string]*gorm.DB
	ring    *hashring.HashRing
}

func NewpostgresClients(clients []postgresClient) *postgresClients {
	hosts := make([]string, len(clients))
	clientsMap := map[string]*gorm.DB{}

	for i, c := range clients {
		hosts[i] = c.host
		clientsMap[c.host] = c.client
	}

	return &postgresClients{
		clients: clientsMap,
		ring:    hashring.New(hosts),
	}
}

func (clients postgresClients) GetShard(key string) *gorm.DB {
	server, _ := clients.ring.GetNode(key)
	return clients.clients[server]
}

func autoMigrate(gormdb *gorm.DB) error {
	return gormdb.AutoMigrate(pgmodels.MigrateModels...)
}

func ConnectPostgres(cfg *config.Config) (IPostgresClients, error) {
	result := make([]postgresClient, 0)

	for _, ip := range cfg.PostgresIPs {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			ip, cfg.PostgresUser,
			cfg.PostgresPassword, cfg.PostgresDB,
			cfg.PostgresPort,
		)

		postgresConn := postgres.Open(dsn)
		gormDB, err := gorm.Open(postgresConn)

		if err != nil {
			return nil, err
		}

		if cfg.PostgresAutoMigrate {
			err := autoMigrate(gormDB)
			if err != nil {
				return nil, err
			}
		}

		client := postgresClient{
			client: gormDB,
			host:   fmt.Sprintf("%s:%d", ip, cfg.PostgresPort),
		}
		result = append(result, client)
	}

	return NewpostgresClients(result), nil
}
