package server

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/contract"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate"
	golog "log"
)

type (
	handler struct {
		cfg         *config.Config
		userService contract.UserService
		logger      log.Logger
		translator  translate.Translator
	}

	HandlerFields struct {
		Cfg         *config.Config
		UserService contract.UserService
		Logger      log.Logger
		Translator  translate.Translator
	}
)

func NewHttpHandler(h *HandlerFields) *handler {
	if h.Cfg == nil {
		golog.Fatal("handler config is nil")
	}

	if h.UserService == nil {
		golog.Fatal("handler user service is nil")
	}

	if h.Logger == nil {
		golog.Fatal("handler logger is nil")
	}

	if h.Translator == nil {
		golog.Fatal("handler translator is nil")
	}

	return &handler{
		cfg:         h.Cfg,
		userService: h.UserService,
		logger:      h.Logger,
		translator:  h.Translator,
	}
}
