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

// Organization ...
type Organization struct {
	BaseModel
	Name    string `json:"name"`
	Address string `json:"address"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
}

// User ...
type User struct {
	BaseModel
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Mobile         string `json:"mobile"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	OrganizationID string `json:"organizationID"`
}
