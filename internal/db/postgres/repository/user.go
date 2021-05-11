package repository

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"github.com/kianooshaz/bookstore-api/pkg/errors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

func (r *Repository) GetUserByID(id uint) (*models.User, error) {
	user := new(models.User)

	if err := r.DB.Model(&models.User{}).Where("id = ?", id).First(user).Error; err != nil {
		r.Logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "GetUserByID",
			Params:   map[string]interface{}{"user_id": id},
			Message:  err.Error(),
		})

		if IsErrorNotFound(err) {
			return nil, errors.New(errors.KindNotFound, messages.UserNotFound)
		}

		return nil, errors.New(errors.KindUnexpected, messages.DBError)
	}

	return user, nil
}

func (r *Repository) GetUserByUsername(username string) (*models.User, error) {
	user := new(models.User)

	if err := r.DB.Model(&models.User{}).Where("username = ?", username).First(user).Error; err != nil {
		r.Logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "GetUserByUsername",
			Params:   map[string]interface{}{"username": username},
			Message:  err.Error(),
		})

		if IsErrorNotFound(err) {
			return nil, errors.New(errors.KindNotFound, messages.UserNotFound)
		}

		return nil, errors.New(errors.KindUnexpected, messages.DBError)
	}

	return user, nil
}

func (r *Repository) UpdateUser(user *models.User) error {
	if err := r.DB.Model(&models.User{}).Save(user).Error; err != nil {
		r.Logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "UpdateUser",
			Params:   map[string]interface{}{"user": user},
			Message:  err.Error(),
		})

		return errors.New(errors.KindUnexpected, messages.DBError)
	}

	return nil
}

func (r *Repository) DeleteUserByID(id uint) error {
	res := r.DB.Model(models.User{}).Where("id = ?", id).Delete(&models.User{})

	if err := res.Error; err != nil {
		r.Logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "DeleteUserByID",
			Params:   map[string]interface{}{"user_id": id},
			Message:  err.Error(),
		})

		return errors.New(errors.KindUnexpected, messages.DBError)
	}

	if res.RowsAffected != 1 {
		r.Logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "DeleteUserByID",
			Params:   map[string]interface{}{"user_id": id},
			Message:  r.Translator.TranslateEn(messages.UserNotFound),
		})

		return errors.New(errors.KindNotFound, messages.UserNotFound)
	}

	return nil
}

func (r *Repository) AddUser(user *models.User) (*models.User, *models.Wallet, error) {
	tx := r.DB.Begin()

	res := tx.Model(&models.User{}).Create(user)
	if err := res.Error; err != nil {
		tx.Rollback()

		r.Logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "AddUser",
			Params:   map[string]interface{}{"user": user},
			Message:  err.Error(),
		})

		return nil, nil, errors.New(errors.KindUnexpected, messages.DBError)
	}

	wallet := &models.Wallet{
		UserID:  user.ID,
		Balance: 0,
		Status:  types.WalletOpen,
	}

	res = tx.Model(&models.Wallet{}).Create(wallet)
	if err := res.Error; err != nil {
		tx.Rollback()

		r.Logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "AddUser",
			Params:   map[string]interface{}{"wallet": wallet},
			Message:  err.Error(),
		})

		return nil, nil, errors.New(errors.KindUnexpected, messages.DBError)
	}

	if err := tx.Commit().Error; err != nil {
		r.Logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "AddUser",
			Params:   map[string]interface{}{"user": user, "wallet": wallet},
			Message:  err.Error(),
		})

		return nil, nil, errors.New(errors.KindUnexpected, messages.DBError)
	}

	return user, wallet, nil
}
