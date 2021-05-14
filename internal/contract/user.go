package contract

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/params"
)

type (
	UserService interface {
		CreateUser(req *params.CreateUserRequest) (*models.User, error)
		GetUserByID(userID uint) (*models.User, error)
		GetUserByUsername(username string) (*models.User, error)
		UpdateUser(req *params.UpdateUserRequest) (*models.User, error)
		DeleteUser(userID uint) error
	}
)
