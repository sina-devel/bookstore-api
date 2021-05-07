package postgres

import (
	"fmt"
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/contract"
	"github.com/kianooshaz/bookstore-api/internal/db/postgres/repository"
	"github.com/kianooshaz/bookstore-api/pkg/errors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	db         *gorm.DB
	cfg        *config.Config
	translator translate.Translator
	logger     log.Logger
}

func New(cfg *config.Config, translator translate.Translator, logger log.Logger) (contract.Repository, error) {
	db := &database{
		cfg:        cfg,
		translator: translator,
		logger:     logger,
	}

	if err := db.connect(); err != nil {
		return nil, err
	}

	if db.cfg.Database.Migration {
		if err := db.migration(); err != nil {
			return nil, err
		}
	}

	return &repository.Repository{
		DB:         db.db,
		Cfg:        db.cfg,
		Translator: db.translator,
		Logger:     db.logger,
	}, nil
}

func (d *database) connect() error {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		d.cfg.Database.Host,
		d.cfg.Database.Username,
		d.cfg.Database.Password,
		d.cfg.Database.DBName,
		d.cfg.Database.Port,
		d.cfg.Database.SSLMode,
		d.cfg.Database.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		d.logger.Error(&log.Field{
			Section:  "postgres.postgres",
			Function: "connect",
			Message:  err.Error(),
		})

		return errors.New(errors.KindUnexpected, messages.DBError)
	}

	d.db = db

	return nil
}
