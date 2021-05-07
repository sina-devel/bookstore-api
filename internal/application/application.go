package application

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/db/postgres"
	"github.com/kianooshaz/bookstore-api/pkg/log/logrus"
	"github.com/kianooshaz/bookstore-api/pkg/translate/i18n"
)

func Run(cfg *config.Config) error {
	logger, err := logrus.New(&logrus.Option{
		Path:         cfg.Log.InternalPath,
		Pattern:      cfg.Log.FilenamePattern,
		MaxAge:       cfg.Log.MaxAge,
		RotationTime: cfg.Log.RotationTime,
		RotationSize: cfg.Log.MaxSize,
	})
	if err != nil {
		return err
	}

	translator, err := i18n.New(cfg.I18n.BundlePath)
	if err != nil {
		return err
	}

	repository, err := postgres.New(cfg, translator, logger)

	_ = repository

	return nil
}
