package auth

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/contract"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate"
)

type service struct {
	cfg        *config.Auth
	authRepo   contract.AuthRepository
	logger     log.Logger
	translator translate.Translator
}

func New(cfg config.Auth, authRepo contract.AuthRepository, logger log.Logger, translator translate.Translator) contract.AuthService {
	return service{
		cfg:        &cfg,
		authRepo:   authRepo,
		logger:     logger,
		translator: translator,
	}
}
