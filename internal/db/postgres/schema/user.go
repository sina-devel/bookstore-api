package schema

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
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

func (u *User) ConvertModel() *models.User {
	return &models.User{
		ID:                    u.ID,
		Username:              u.Username,
		FirstName:             u.FirstName,
		LastName:              u.LastName,
		Email:                 u.Email,
		IsEmailVerified:       u.IsEmailVerified,
		PhoneNumber:           u.PhoneNumber,
		IsPhoneNumberVerified: u.IsPhoneNumberVerified,
		Gender:                u.Gender,
		Role:                  u.Role,
		Avatar:                u.Avatar,
		Wallet: models.Wallet{
			ID:      u.Wallet.ID,
			UserID:  u.Wallet.UserID,
			Balance: u.Wallet.Balance,
			Status:  u.Wallet.Status,
		},
	}
}

func ConvertUser(user *models.User) *User {
	return &User{
		Model:                 gorm.Model{ID: user.ID},
		Username:              user.Username,
		FirstName:             user.FirstName,
		LastName:              user.LastName,
		Email:                 user.Email,
		IsEmailVerified:       user.IsEmailVerified,
		PhoneNumber:           user.PhoneNumber,
		IsPhoneNumberVerified: user.IsPhoneNumberVerified,
		Gender:                user.Gender,
		Role:                  user.Role,
		Avatar:                user.Avatar,
		Wallet: Wallet{
			Model:   gorm.Model{ID: user.Wallet.UserID},
			UserID:  user.Wallet.UserID,
			Balance: user.Wallet.Balance,
			Status:  user.Wallet.Status,
		},
	}
}
