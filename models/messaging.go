package models

import (
	"time"

	"gorm.io/gorm"
)

// Inbox ...
type Inbox struct {
	ID          string         `json:"id"`
	From        string         `json:"from"`
	To          string         `json:"to"` // can be shortCode or SenderID
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
	From        string         `json:"from"` // can be shortCode or SenderID
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
