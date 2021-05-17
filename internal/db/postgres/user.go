package postgres

import (
	"github.com/kianooshaz/bookstore-api/internal/db/postgres/schema"
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

//CreateUser creates an user and wallet for this user on the database users and wallets tables
func (r *repository) CreateUser(user *models.User) (*models.User, error) {
	u := schema.ConvertUser(user)

	if err := r.db.Create(u).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "CreateUser",
			Params:   map[string]interface{}{"user": u},
			Message:  err.Error(),
		})

		return nil, derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	return u.ConvertModel(), nil
}

func (r *repository) GetUserByID(userID uint) (*models.User, error) {
	user := new(schema.User)

	if err := r.db.Model(&schema.User{}).Where("id = ?", userID).Preload("Wallet").First(user).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "GetUserByID",
			Params:   map[string]interface{}{"user_id": userID},
			Message:  err.Error(),
		})

		if isErrorNotFound(err) {
			return nil, derrors.New(derrors.KindNotFound, messages.UserNotFound)
		}

		return nil, derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	return user.ConvertModel(), nil
}

func (r *repository) GetUserByUsername(username string) (*models.User, error) {
	user := new(schema.User)

	if err := r.db.Model(&schema.User{}).Where("username = ?", username).Preload("Wallet").First(user).Error; err != nil {
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

	return user.ConvertModel(), nil
}

func (r *repository) UpdateUser(user *models.User) (*models.User, error) {
	u := schema.ConvertUser(user)

	if err := r.db.Model(&schema.User{}).First(&schema.User{}, u.ID).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "UpdateUser",
			Params:   map[string]interface{}{"user": u},
			Message:  err.Error(),
		})

		if isErrorNotFound(err) {
			return nil, derrors.New(derrors.KindNotFound, messages.UserNotFound)
		}

		return nil, derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	if err := r.db.Model(&schema.User{}).Where("id = ?", u.ID).Save(u).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "UpdateUser",
			Params:   map[string]interface{}{"user": u},
			Message:  err.Error(),
		})

		return nil, derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	return u.ConvertModel(), nil
}

func (r *repository) DeleteUser(user *models.User) error {
	u := schema.ConvertUser(user)

	res := r.db.Select("Wallet").Where("id", u.ID).Delete(u)
	if err := res.Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "DeleteUser",
			Params:   map[string]interface{}{"user": u},
			Message:  err.Error(),
		})

		return derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	if res.RowsAffected != 1 {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "DeleteUser",
			Params:   map[string]interface{}{"user": u},
			Message:  r.translator.TranslateEn(messages.UserNotFound),
		})

		return derrors.New(derrors.KindNotFound, messages.UserNotFound)
	}

	return nil
}
