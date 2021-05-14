package user

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/contract"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate"
)

type service struct {
	cfg        *config.Config
	userRepo   contract.UserRepository
	logger     log.Logger
	translator translate.Translator
}

func New(cfg *config.Config, mainRepo contract.MainRepository, logger log.Logger, translator translate.Translator) contract.UserService {
	return &service{
		cfg:        cfg,
		userRepo:   mainRepo,
		logger:     logger,
		translator: translator,
	}
}
