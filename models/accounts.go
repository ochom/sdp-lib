package models

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
