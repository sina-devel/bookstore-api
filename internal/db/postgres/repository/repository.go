package repository

import (
	"errors"
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate"
	"gorm.io/gorm"
)

type Repository struct {
	DB         *gorm.DB
	Cfg        *config.Config
	Translator translate.Translator
	Logger     log.Logger
}

func IsErrorNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
