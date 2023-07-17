package pgmodels

import "gorm.io/gorm"

type TimingAt struct {
	CreatedAt uint64         `gorm:"column:created_at;type:timestamp with time zone;not null;now" json:"created_at"`
	UpdatedAt uint64         `gorm:"column:updated_at;type:timestamp with time zone;not null;now" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone" json:"deleted_at"`
}

var MigrateModels = []interface{}{
	PostgresLogMsg{},
}
