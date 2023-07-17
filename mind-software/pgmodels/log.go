package pgmodels

type PostgresLogMsg struct {
	ID        uint64 `gorm:"column:id;type:serial;primaryKey;autoincrement" json:"id"`
	ObjID     uint64 `gorm:"column:obj_id;type:bigint;primaryKey;not null" json:"obj_id"`
	Timestamp uint64 `gorm:"column:timestamp;type:bigint;not null" json:"timestamp"`
	Log       string `gorm:"column:log;type:text;not null" json:"log"`

	TimingAt
}
