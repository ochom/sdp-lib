package models

import (
	"time"

	"gorm.io/gorm"
)

// ContactGroup ...
type ContactGroup struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	OrganizationID string         `json:"organizationID"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
	TotalContacts  int64          `json:"totalContacts" gorm:"-"`
}

// Contact ...
type Contact struct {
	ID        string         `json:"id"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Mobile    string         `json:"mobile"`
	Email     string         `json:"email"`
	GroupID   string         `json:"groupID"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
