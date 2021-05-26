package auth

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/contract"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate"
)

type service struct {
	cfg        *config.Auth
	logger     log.Logger
	translator translate.Translator
}

func New(cfg config.Auth, logger log.Logger, translator translate.Translator) contract.AuthService {
	return service{
		cfg:        &cfg,
		logger:     logger,
		translator: translator,
	}
}
