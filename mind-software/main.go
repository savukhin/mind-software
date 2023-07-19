package main

import (
	"mind-software/config"
	"mind-software/connection"
	"mind-software/handlers"
	"net/http"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/nats-io/nats.go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		logger = zap.NewNop()
	}
	defer logger.Sync()

	cfg := &config.Config{}
	err = envconfig.Process("", cfg)
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}

	postgresClients, err := connection.ConnectPostgres(cfg)
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	nc, err := nats.Connect(cfg.NatsUrl,
		nats.ClosedHandler(func(c *nats.Conn) {
			wg.Done()
		}),
	)
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}

	ec.QueueSubscribe(cfg.NatsLogsTopic, "workers", handlers.AddLog(postgresClients, logger))

	ec.QueueSubscribe(cfg.NatsQueryTopic, "workers", handlers.GetLogs(postgresClients, logger, ec, cfg.NatsResponseTopic))

	// wg.Wait()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
