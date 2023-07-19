package handlers

import (
	"errors"
	"fmt"
	"mind-software/connection"
	"mind-software/pgmodels"
	"mind-software/query"
	"time"

	"go.uber.org/zap"
)

type NatsLogMsg struct {
	ObjId uint64 `json:"id"`
	Log   string `json:"log"`
}

func AddLog(clients connection.IPostgresClients, logger *zap.Logger) func(msg *NatsLogMsg) error {
	return func(msg *NatsLogMsg) error {
		if msg == nil {
			logger.Error("AddLog received nil message")
			return errors.New("nil message")
		}
		gormDB := clients.GetShard(fmt.Sprintf("%d", msg.ObjId))
		table := query.Use(gormDB).PostgresLogMsg

		model := &pgmodels.PostgresLogMsg{
			ObjID:     msg.ObjId,
			Timestamp: uint64(time.Now().UnixMilli()),
			Log:       msg.Log,
		}

		err := table.Create(model)
		if err != nil {
			logger.Error("Table create error " + err.Error())
			return err
		}
		return nil
	}
}
