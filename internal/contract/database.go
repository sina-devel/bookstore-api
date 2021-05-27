package contract

import "github.com/kianooshaz/bookstore-api/internal/models"

type (
	MainRepository interface {
		UserRepository
		AuthRepository
	}

	UserRepository interface {
		CreateUser(user *models.User) (*models.User, error)
		GetUserByID(id uint) (*models.User, error)
		GetUserByUsername(username string) (*models.User, error)
		UpdateUser(user *models.User) (*models.User, error)
		DeleteUser(user *models.User) error
		IsUsernameExist(username string) (bool, error)
	}

	AuthRepository interface {
		CreateToken(token string, userID uint) error
		TokenIsExistWithUserID(token string, userID uint) (bool, error)
	}
)
