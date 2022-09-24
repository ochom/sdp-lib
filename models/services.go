package models

// Service ...
type Service struct {
	BaseModel
	Name           string `json:"name"`
	OfferName      string `json:"offerName"`
	OfferCode      string `json:"offerCode"`
	Description    string `json:"description"`
	OrganizationID string `json:"organizationID"`
}

// Subscriber ...
type Subscriber struct {
	BaseModel
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	ServiceID string `json:"organizationID"`
}
