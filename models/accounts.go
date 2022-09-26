package models

import (
	"time"

	"gorm.io/gorm"
)

// Organization ...
type Organization struct {
	ID         string         `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name"`
	Address    string         `json:"address"`
	Mobile     string         `json:"mobile"`
	Email      string         `json:"email"`
	IsVerified bool           `json:"isVerified"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt"`
}

// User ...
type User struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	Mobile         string         `json:"mobile"`
	Email          string         `json:"email"`
	Password       string         `json:"-"`
	OTP            string         `json:"-"`
	OrganizationID string         `json:"organizationID"`
	IsVerified     bool           `json:"isVerified"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
}
