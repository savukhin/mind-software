package handlers

import (
	"errors"
	"fmt"
	"mind-software/connection"
	"mind-software/pgmodels"
	"mind-software/query"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type NatsGetLogsMsg struct {
	ObjId uint64 `json:"id"`
}

func getLogs(msg *NatsGetLogsMsg, clients connection.IPostgresClients, logger *zap.Logger) ([]*pgmodels.PostgresLogMsg, error) {
	if msg == nil {
		return nil, errors.New("received nil message")
	}
	gormDB := clients.GetShard(fmt.Sprintf("%d", msg.ObjId))
	table := query.Use(gormDB).PostgresLogMsg

	logs, err := table.Where(table.ObjID.Eq(msg.ObjId)).Order(table.Timestamp).Find()
	return logs, err
}

func GetLogs(clients connection.IPostgresClients, logger *zap.Logger, ec *nats.EncodedConn, natsResponseTopic string) func(msg *NatsGetLogsMsg) {
	return func(msg *NatsGetLogsMsg) {
		logs, err := getLogs(msg, clients, logger)
		if err != nil {
			logger.Error(err.Error())
			return
		}

		err = ec.Publish(natsResponseTopic, logs)
		if err != nil {
			logger.Error("Response NATS err " + err.Error())
		}
	}
}
