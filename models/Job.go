package models

import (
	"time"
)

type Job struct {
	// gorm.Model
	ID          int `gorm:"uniqueIndex:compositeindex;type:text;not null"`
	StoreId     string
	CreatedAt   time.Time
	CompletedAt time.Time
	Status      string
}
