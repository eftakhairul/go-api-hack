package models

import (
	"time"
)

// RequestBody interface for any model that used for validation
type RequestBody interface {
	Validate() error
}

// Model definition same as gorm.Model, but including column and json tags
type Model struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
