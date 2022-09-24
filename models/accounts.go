package models

import (
	"time"

	"gorm.io/gorm"
)

// Organization ...
type Organization struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	Mobile    string         `json:"mobile"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

// User ...
type User struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	Mobile         string         `json:"mobile"`
	Email          string         `json:"email"`
	Password       string         `json:"password"`
	OrganizationID string         `json:"organizationID"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
}
