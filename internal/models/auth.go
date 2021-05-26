package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kianooshaz/bookstore-api/internal/models/types"
)

type Claim struct {
	ID          uint       `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Role        types.Role `json:"role"`
	jwt.StandardClaims
}
