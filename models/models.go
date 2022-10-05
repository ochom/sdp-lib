package models

// AllModels ...
func AllModels() []interface{} {
	return []interface{}{
		&Organization{},
		&Shortcode{},
		&User{},
		&ContactGroup{},
		&Contact{},
		&Subscriber{},
	}
}
