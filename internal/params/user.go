package params

import "github.com/kianooshaz/bookstore-api/internal/models/types"

type (
	CreateUserRequest struct {
		Username              string       `json:"username"`
		Password              string       `json:"password"`
		FirstName             string       `json:"first_name"`
		LastName              string       `json:"last_name"`
		Email                 string       `json:"email"`
		IsEmailVerified       bool         `json:"is_email_verified"`
		PhoneNumber           string       `json:"phone_number"`
		IsPhoneNumberVerified bool         `json:"is_phone_number_verified"`
		Gender                types.Gender `json:"gender"`
		Role                  types.Role   `json:"role"`
	}

	UpdateUserRequest struct {
		ID          uint         `json:"id"`
		FirstName   string       `json:"first_name"`
		LastName    string       `json:"last_name"`
		Email       string       `json:"email"`
		PhoneNumber string       `json:"phone_number"`
		Gender      types.Gender `json:"gender"`
		Avatar      string       `json:"avatar"`
	}

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
