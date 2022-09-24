package models

// ContactGroup ...
type ContactGroup struct {
	BaseModel
	Name           string `json:"name"`
	Description    string `json:"description"`
	OrganizationID string `json:"organizationID"`
}

// Contact ...
type Contact struct {
	BaseModel
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	GroupID   string `json:"groupID"`
}
