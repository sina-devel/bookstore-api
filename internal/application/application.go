package application

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/db/postgres"
	"github.com/kianooshaz/bookstore-api/internal/http/server"
	"github.com/kianooshaz/bookstore-api/internal/service/user"
	"github.com/kianooshaz/bookstore-api/pkg/log/logrus"
	"github.com/kianooshaz/bookstore-api/pkg/translate/i18n"
)

func Run(cfg *config.Config) error {
	logger, err := logrus.New(&logrus.Option{
		Path:         cfg.Logger.InternalPath,
		Pattern:      cfg.Logger.FilenamePattern,
		MaxAge:       cfg.Logger.MaxAge,
		RotationTime: cfg.Logger.RotationTime,
		RotationSize: cfg.Logger.MaxSize,
	})
	if err != nil {
		return err
	}

	translator, err := i18n.New(cfg.I18n.BundlePath)
	if err != nil {
		return err
	}

	mainRepository, err := postgres.New(cfg.Database.Postgres, translator, logger)
	if err != nil {
		return err
	}

	userService := user.New(cfg.User, mainRepository, logger, translator)

	handler := server.NewHttpHandler(&server.HandlerFields{
		Cfg:         cfg,
		UserService: userService,
		Logger:      logger,
		Translator:  translator,
	})
	httpServer := server.NewHttpServer(handler)

	return httpServer.Start(8083)
}
