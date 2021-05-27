package postgres

import (
	"github.com/kianooshaz/bookstore-api/internal/db/postgres/schema"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

func (r *repository) CreateToken(token string, userID uint) error {
	t := &schema.Token{
		Value:  token,
		UserID: userID,
	}

	if err := r.db.Model(&schema.Token{}).Where("user_id = ?", userID).Delete(&schema.Token{}).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.auth",
			Function: "CreateToken",
			Params:   map[string]interface{}{"user_id": userID},
			Message:  err.Error(),
		})

		return derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	if err := r.db.Model(&schema.Token{}).Create(t).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.auth",
			Function: "CreateToken",
			Params:   map[string]interface{}{"user_id": userID},
			Message:  err.Error(),
		})

		return derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	return nil
}

func (r *repository) TokenIsExistWithUserID(token string, userID uint) (bool, error) {
	t := &schema.Token{}
	if err := r.db.Model(&schema.Token{}).Where("value = ? and user_id = ?", token, userID).First(&t).Error; err != nil {

		if isErrorNotFound(err) {
			return false, nil
		}

		r.logger.Error(&log.Field{
			Section:  "repository.auth",
			Function: "TokenIsExistWithUserID",
			Message:  err.Error(),
		})

		return false, derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	return true, nil
}
