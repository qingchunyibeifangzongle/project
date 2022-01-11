package model

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	CreatedAt time.Time      `gorm:"type:timestamp;column:created_at;index" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;column:updated_at;index" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;column:deleted_at" json:"deleted_at"`
}
