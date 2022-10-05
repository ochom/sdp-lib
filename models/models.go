package models

// AllModels ...
func AllModels() []interface{} {
	return []interface{}{
		&Organization{},
		&Offer{},
		&User{},
		&ContactGroup{},
		&Contact{},
		&Subscriber{},
	}
}
