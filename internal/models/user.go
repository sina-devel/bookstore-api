package models

import (
	"github.com/kianooshaz/bookstore-api/internal/models/types"
)

type (
	User struct {
		ID                    uint         `json:"id"`
		Username              string       `json:"username"`
		Password              string       `json:"-"`
		FirstName             string       `json:"first_name"`
		LastName              string       `json:"last_name"`
		Email                 string       `json:"email"`
		IsEmailVerified       bool         `json:"is_email_verified"`
		PhoneNumber           string       `json:"phone_number"`
		IsPhoneNumberVerified bool         `json:"is_phone_number_verified"`
		Gender                types.Gender `json:"gender"`
		Role                  types.Role   `json:"role"`
		Avatar                string       `json:"avatar"`
		Wallet                Wallet       `json:"wallet"`
	}
)
