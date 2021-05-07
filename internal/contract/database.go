package contract

import "github.com/kianooshaz/bookstore-api/internal/models"

type (
	Repository interface {
		UserRepository
	}

	UserRepository interface {
		CreateUser(user *models.User) error
	}
)
