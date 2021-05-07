package repository

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/pkg/errors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

func (r *Repository) CreateUser(user *models.User) error {
	if err := r.DB.Create(user).Error; err != nil {
		r.Logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "CreateUser",
			Params:   user,
			Message:  err.Error(),
		})

		return errors.New(errors.KindUnexpected, messages.DBError)
	}

	return nil
}
