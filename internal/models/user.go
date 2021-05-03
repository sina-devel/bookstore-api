package models

import "github.com/kianooshaz/bookstore-api/internal/models/types"

type (
	User struct {
		Base
		Username              string `gorm:"unique;not null"`
		FirstName             string
		LastName              string
		Email                 string
		IsEmailVerified       bool
		PhoneNumber           string
		IsPhoneNumberVerified bool
		Gender                types.Gender
		Role                  types.Role
		Avatar                string
		Token                 string
	}
)
