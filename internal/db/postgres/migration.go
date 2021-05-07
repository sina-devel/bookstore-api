package postgres

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/pkg/errors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

func (d *database) migration() error {
	if err := d.db.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Wallet{},
		&models.Category{},
		&models.Comment{},
		&models.Picture{},
	); err != nil {
		d.logger.Error(&log.Field{
			Section:  "postgres.migration",
			Function: "migration",
			Message:  err.Error(),
		})

		return errors.New(errors.KindUnexpected, messages.DBError)
	}

	return nil
}
