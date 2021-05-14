package postgres

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/contract"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate"
	"gorm.io/gorm"
)

type repository struct {
	db         *gorm.DB
	cfg        *config.Postgres
	translator translate.Translator
	logger     log.Logger
}

func New(cfg *config.Postgres, translator translate.Translator, logger log.Logger) (contract.MainRepository, error) {
	repo := &repository{
		cfg:        cfg,
		translator: translator,
		logger:     logger,
	}

	if err := repo.connect(); err != nil {
		return nil, err
	}

	if cfg.Migration {
		if err := repo.migration(); err != nil {
			return nil, err
		}
	}

	return repo, nil
}
