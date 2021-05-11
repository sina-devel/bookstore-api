package contract

import "github.com/kianooshaz/bookstore-api/internal/models"

type (
	MainRepository interface {
		UserRepository
	}

	UserRepository interface {
		GetUserByID(id uint) (*models.User, error)
		GetUserByUsername(username string) (*models.User, error)
		UpdateUser(user *models.User) error
		DeleteUserByID(id uint) error
		AddUser(user *models.User) (*models.User, *models.Wallet, error)
	}
)
