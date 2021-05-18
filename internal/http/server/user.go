package server

import (
	"github.com/kianooshaz/bookstore-api/internal/params"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) createUser(c echo.Context) error {
	lang := getLanguage(c)
	req := new(params.CreateUserRequest)

	if err := c.Bind(req); err != nil {
		h.logger.Error(&log.Field{
			Section:  "server",
			Function: "createUser",
			Message:  h.translator.TranslateEn(err.Error()),
		})

		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(lang, messages.ParseQueryError),
		}
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		message, code := derrors.HttpError(err)

		h.logger.Error(&log.Field{
			Section:  "server",
			Function: "createUser",
			Params:   map[string]interface{}{"req": req},
			Message:  h.translator.TranslateEn(message),
		})

		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(lang, message),
		}
	}

	return c.JSON(http.StatusOK, user)
}
