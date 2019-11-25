package models

import (
	"time"
)

type RequestBody interface {
	Validate() error
}

// Model definition same as gorm.Model, but including column and json tags
type Model struct {
	ID        uint       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
