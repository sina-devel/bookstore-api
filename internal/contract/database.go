package contract

import "github.com/kianooshaz/bookstore-api/internal/models"

type (
	MainRepository interface {
		UserRepository
	}

	UserRepository interface {
		CreateUser(user *models.User) (*models.User, error)
		GetUserByID(id uint) (*models.User, error)
		GetUserByUsername(username string) (*models.User, error)
		UpdateUser(user *models.User) (*models.User, error)
		DeleteUser(user *models.User) error
	}
)
