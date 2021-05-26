package user

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/contract"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate"
)

type service struct {
	userCfg     *config.User
	userRepo    contract.UserRepository
	authService contract.AuthService
	logger      log.Logger
	translator  translate.Translator
}

func New(userCfg config.User, mainRepo contract.MainRepository, authService contract.AuthService, logger log.Logger, translator translate.Translator) contract.UserService {
	return &service{
		userCfg:     &userCfg,
		userRepo:    mainRepo,
		authService: authService,
		logger:      logger,
		translator:  translator,
	}
}
