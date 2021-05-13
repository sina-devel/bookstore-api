package postgres

import (
	"fmt"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (r *repository) connect() error {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		r.cfg.Database.Postgres.Host,
		r.cfg.Database.Postgres.Username,
		r.cfg.Database.Postgres.Password,
		r.cfg.Database.Postgres.DBName,
		r.cfg.Database.Postgres.Port,
		r.cfg.Database.Postgres.SSLMode,
		r.cfg.Database.Postgres.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		r.logger.Error(&log.Field{
			Section:  "postgres.postgres",
			Function: "connect",
			Message:  err.Error(),
		})

		return derrors.New(derrors.KindUnexpected, messages.DBError)
	}

	r.db = db

	return nil
}
