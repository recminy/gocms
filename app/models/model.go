package models

import (
	"database/sql"
	"time"
)

type PrimaryKeyField struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`
}

type TimestampField struct {
	DeletedAt sql.NullTime `json:"deleted_at"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
