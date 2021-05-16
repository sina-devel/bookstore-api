package user

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/params"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/hash"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

func (s *service) CreateUser(req *params.CreateUserRequest) (*models.User, error) {
	if err := s.validateUsername(req.Username); err != nil {
		return nil, err
	}

	if err := s.validatePassword(req.Password); err != nil {
		return nil, err
	}

	password, err := hash.Password(req.Password)
	if err != nil {
		s.logger.Error(&log.Field{
			Section:  "service.user",
			Function: "CreateUser",
			Params:   map[string]interface{}{"password": req.Password},
			Message:  err.Error(),
		})

		return nil, derrors.New(derrors.KindUnexpected, messages.GeneralError)
	}

	user := &models.User{
		Username:              req.Username,
		Password:              password,
		FirstName:             req.FirstName,
		LastName:              req.LastName,
		Email:                 req.Email,
		IsEmailVerified:       req.IsEmailVerified,
		PhoneNumber:           req.PhoneNumber,
		IsPhoneNumberVerified: req.IsPhoneNumberVerified,
		Gender:                req.Gender,
		Role:                  req.Role,
	}

	user, err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByID(userID uint) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByUsername(username string) (*models.User, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) UpdateUser(req *params.UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(req.ID)
	if err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.IsEmailVerified = req.IsEmailVerified
	user.PhoneNumber = req.PhoneNumber
	user.IsPhoneNumberVerified = req.IsPhoneNumberVerified
	user.Gender = req.Gender
	user.Role = req.Role
	user.Avatar = req.Avatar

	user, err = s.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *service) DeleteUser(userID uint) error {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}

	if err := s.userRepo.DeleteUser(user); err != nil {
		return err
	}

	return nil

}
