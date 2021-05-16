package user

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/contract"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate"
)

type service struct {
	userCfg    *config.User
	userRepo   contract.UserRepository
	logger     log.Logger
	translator translate.Translator
}

func New(userCfg config.User, mainRepo contract.MainRepository, logger log.Logger, translator translate.Translator) contract.UserService {
	return &service{
		userCfg:    &userCfg,
		userRepo:   mainRepo,
		logger:     logger,
		translator: translator,
	}
}
