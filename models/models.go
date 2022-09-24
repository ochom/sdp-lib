package models

// AllModels ...
func AllModels() []interface{} {
	return []interface{}{
		&Organization{},
		&User{},
		&ContactGroup{},
		&Contact{},
		&Subscriber{},
	}
}
