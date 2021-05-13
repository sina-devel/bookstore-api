package models

import (
	"github.com/kianooshaz/bookstore-api/internal/models/types"
)

type (
	User struct {
		ID                    uint
		Username              string
		FirstName             string
		LastName              string
		Email                 string
		IsEmailVerified       bool
		PhoneNumber           string
		IsPhoneNumberVerified bool
		Gender                types.Gender
		Role                  types.Role
		Avatar                string
		Wallet                Wallet
	}
)
