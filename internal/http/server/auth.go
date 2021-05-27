package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kianooshaz/bookstore-api/internal/params"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *handler) login(c echo.Context) error {
	lang := getLanguage(c)

	req := new(params.LoginRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error(&log.Field{
			Section:  "http.server",
			Function: "login",
			Message:  err.Error(),
		})

		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(lang, messages.ParseQueryError),
		}
	}

	tokens, err := h.userService.Login(req)
	if err != nil {
		message, code := derrors.HttpError(err)
		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(lang, message),
		}
	}

	return c.JSON(http.StatusOK, tokens)
}

func (h *handler) refreshToken(c echo.Context) error {
	lang := getLanguage(c)

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		h.logger.Error(&log.Field{
			Section:  "server.auth",
			Function: "refreshToken",
			Message:  h.translator.TranslateEn(messages.InvalidToken),
		})

		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: h.translator.Translate(lang, messages.InvalidToken),
		}
	}

	userIdString := c.Param("id")
	userID, err := strconv.Atoi(userIdString)
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.auth",
			Function: "refreshToken",
			Params:   map[string]interface{}{"user_id_string": userIdString},
			Message:  err.Error(),
		})

		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(lang, messages.ParseQueryError),
		}
	}

	res, err := h.userService.RefreshToken(token.Raw, uint(userID))
	if err != nil {
		message, code := derrors.HttpError(err)

		return &echo.HTTPError{
			Code:    code,
			Message: message,
		}
	}

	return c.JSON(http.StatusOK, res)
}
