package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel ...
type BaseModel struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

// AllModels ...
func AllModels() []interface{} {
	return []interface{}{
		&Organization{},
		&User{},
		&ContactGroup{},
		&Contact{},
		&Subscriber{},
	}
}
