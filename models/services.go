package models

import (
	"time"

	"gorm.io/gorm"
)

// Service ...
type Service struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name"`
	OfferName      string         `json:"offerName"`
	OfferCode      string         `json:"offerCode"`
	Description    string         `json:"description"`
	OrganizationID string         `json:"organizationID"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
}

// Subscriber ...
type Subscriber struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Mobile    string         `json:"mobile"`
	Email     string         `json:"email"`
	ServiceID string         `json:"organizationID"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
