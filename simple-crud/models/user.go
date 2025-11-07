package models

type User struct {
	FirstName string `validate:"required" json:"first_name"`
	LastName  string `validate:"required" json:"last_name"`
	Biography string `validate:"required" json:"biography"`
}

type ID string
