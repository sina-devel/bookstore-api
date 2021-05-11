package postgres

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

func (r *repository) GetUserByID(id uint) (*models.User, error) {
	user := new(models.User)

	if err := r.db.Model(&models.User{}).Where("id = ?", id).First(user).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "GetUserByID",
			Params:   map[string]interface{}{"user_id": id},
			Message:  err.Error(),
		})

		if isErrorNotFound(err) {
			return nil, derrors.New(derrors.KindNotFound, messages.UserNotFound)
		}

		return nil, derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	return user, nil
}

func (r *repository) GetUserByUsername(username string) (*models.User, error) {
	user := new(models.User)

	if err := r.db.Model(&models.User{}).Where("username = ?", username).First(user).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "GetUserByUsername",
			Params:   map[string]interface{}{"username": username},
			Message:  err.Error(),
		})

		if isErrorNotFound(err) {
			return nil, derrors.New(derrors.KindNotFound, messages.UserNotFound)
		}

		return nil, derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	return user, nil
}

func (r *repository) UpdateUser(user *models.User) error {
	if err := r.db.Model(&models.User{}).Save(user).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "UpdateUser",
			Params:   map[string]interface{}{"user": user},
			Message:  err.Error(),
		})

		return derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	return nil
}

func (r *repository) DeleteUserByID(id uint) error {
	res := r.db.Model(&models.User{}).Where("id = ?", id).Delete(&models.User{})

	if err := res.Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "DeleteUserByID",
			Params:   map[string]interface{}{"user_id": id},
			Message:  err.Error(),
		})

		return derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	if res.RowsAffected != 1 {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "DeleteUserByID",
			Params:   map[string]interface{}{"user_id": id},
			Message:  r.translator.TranslateEn(messages.UserNotFound),
		})

		return derrors.New(derrors.KindNotFound, messages.UserNotFound)
	}

	return nil
}

func (r *repository) AddUser(user *models.User) (*models.User, *models.Wallet, error) {
	tx := r.db.Begin()

	res := tx.Model(&models.User{}).Create(user)
	if err := res.Error; err != nil {
		tx.Rollback()

		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "AddUser",
			Params:   map[string]interface{}{"user": user},
			Message:  err.Error(),
		})

		return nil, nil, derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	wallet := &models.Wallet{
		UserID:  user.ID,
		Balance: 0,
		Status:  types.WalletOpen,
	}

	res = tx.Model(&models.Wallet{}).Create(wallet)
	if err := res.Error; err != nil {
		tx.Rollback()

		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "AddUser",
			Params:   map[string]interface{}{"wallet": wallet},
			Message:  err.Error(),
		})

		return nil, nil, derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	if err := tx.Commit().Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "AddUser",
			Params:   map[string]interface{}{"user": user, "wallet": wallet},
			Message:  err.Error(),
		})

		return nil, nil, derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	return user, wallet, nil
}
