package postgres

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

func (r *repository) migration() error {
	if err := r.db.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Wallet{},
		&models.Category{},
		&models.Comment{},
		&models.Picture{},
	); err != nil {
		r.logger.Error(&log.Field{
			Section:  "postgres.migration",
			Function: "migration",
			Message:  err.Error(),
		})

		return derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	return nil
}
