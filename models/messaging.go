package models

// Inbox ...
type Inbox struct {
	BaseModel
	From        string `json:"from"`
	To          string `json:"to"` // can be shortCode or SenderID
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	Attachments string `json:"attachments"`
	RequestID   string `json:"requestID"`
	ResponseID  string `json:"responseID"` // can be LinkID
}

// Outbox ...
type Outbox struct {
	BaseModel
	From        string `json:"from"` // can be shortCode or SenderID
	To          string `json:"to"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	Attachments string `json:"attachments"`
	RequestID   string `json:"requestID"`
	ResponseID  string `json:"responseID"` // can be LinkID
}
