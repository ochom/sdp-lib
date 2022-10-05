package models

import (
	"time"

	"gorm.io/gorm"
)

// Offer ...
type Offer struct {
	ID             string         `json:"id"`
	OrganizationID string         `json:"organizationID"`
	OrgName        string         `json:"orgName"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Shortcode      string         `json:"shortcode"`
	OfferCode      string         `json:"offerCode"`
	CpID           string         `json:"cpID"`
	CpUsername     string         `json:"cpUsername"`
	CpPassword     string         `json:"cpPassword"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
}

// Inbox ...
type Inbox struct {
	ID          string         `json:"id"`
	From        string         `json:"from"`
	To          string         `json:"to"` // can be Offer or SenderID
	Subject     string         `json:"subject"`
	Body        string         `json:"body"`
	Attachments string         `json:"attachments"`
	RequestID   string         `json:"requestID"`
	ResponseID  string         `json:"responseID"` // can be LinkID
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}

// Outbox ...
type Outbox struct {
	ID          string         `json:"id"`
	From        string         `json:"from"` // can be Offer or SenderID
	To          string         `json:"to"`
	Subject     string         `json:"subject"`
	Body        string         `json:"body"`
	Attachments string         `json:"attachments"`
	RequestID   string         `json:"requestID"`
	ResponseID  string         `json:"responseID"` // can be LinkID
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}
