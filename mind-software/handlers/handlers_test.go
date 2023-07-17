package handlers

import (
	"mind-software/pgmodels"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DummyClients struct {
	db *gorm.DB
}

func (c DummyClients) GetShard(key string) *gorm.DB {
	return c.db
}

func GetExecutablePath() string {
	ex, _ := os.Executable()
	return filepath.Dir(ex)
}

func TestGetLogs(t *testing.T) {
	gormDB, err := gorm.Open(sqlite.Open(GetExecutablePath()+"/test.db"), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	gormDB.AutoMigrate(pgmodels.MigrateModels...)

	require.NoError(t, err)
	postgresClients := &DummyClients{db: gormDB}
	logger := zap.NewNop()

	funcAdd := AddLog(postgresClients, logger)

	err = funcAdd(&NatsLogMsg{
		ObjId: 1,
		Log:   "My super Log1",
	})
	require.NoError(t, err)

	err = funcAdd(&NatsLogMsg{
		ObjId: 1,
		Log:   "My super Log2",
	})
	require.NoError(t, err)

	err = funcAdd(nil)
	require.Error(t, err)

	err = funcAdd(&NatsLogMsg{
		ObjId: 2,
		Log:   "My super Log3",
	})
	require.NoError(t, err)

	err = funcAdd(&NatsLogMsg{
		ObjId: 1,
		Log:   "My super Log4",
	})
	require.NoError(t, err)

	logs, err := getLogs(&NatsGetLogsMsg{
		ObjId: 1,
	}, postgresClients, logger)
	require.NoError(t, err)
	require.Len(t, logs, 3)
	require.EqualValues(t, "My super Log1", logs[0].Log)
	require.EqualValues(t, "My super Log2", logs[1].Log)
	require.EqualValues(t, "My super Log4", logs[2].Log)

	logs, err = getLogs(&NatsGetLogsMsg{
		ObjId: 2,
	}, postgresClients, logger)
	require.NoError(t, err)
	require.Len(t, logs, 1)
	require.EqualValues(t, "My super Log3", logs[0].Log)

	logs, err = getLogs(&NatsGetLogsMsg{
		ObjId: 3,
	}, postgresClients, logger)
	require.NoError(t, err)
	require.Len(t, logs, 0)
}
