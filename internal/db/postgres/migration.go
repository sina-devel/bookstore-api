package postgres

import (
	"github.com/kianooshaz/bookstore-api/internal/db/postgres/schema"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

func (r *repository) migration() error {
	if err := r.db.AutoMigrate(
		&schema.User{},
		&schema.Book{},
		&schema.Wallet{},
		&schema.Category{},
		&schema.Comment{},
		&schema.Picture{},
		&schema.Token{},
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
