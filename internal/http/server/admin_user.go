package server

import (
	"github.com/kianooshaz/bookstore-api/internal/params"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *handler) adminCreateUser(c echo.Context) error {
	lang := getLanguage(c)

	req := new(params.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error(&log.Field{
			Section:  "http.server",
			Function: "adminCreateUser",
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

		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(lang, message),
		}
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) adminGetUser(c echo.Context) error {
	lang := getLanguage(c)

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "http.server",
			Function: "adminGetUser",
			Params:   map[string]interface{}{"id": idString},
			Message:  err.Error(),
		})

		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(lang, messages.ParseQueryError),
		}
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		message, code := derrors.HttpError(err)

		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(lang, message),
		}
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) adminUpdateUser(c echo.Context) error {
	lang := getLanguage(c)

	req := new(params.UpdateUserRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error(&log.Field{
			Section:  "http.server",
			Function: "adminUpdateUser",
			Message:  err.Error(),
		})

		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(lang, messages.ParseQueryError),
		}
	}

	user, err := h.userService.UpdateUser(req)
	if err != nil {

		message, code := derrors.HttpError(err)
		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(lang, message),
		}
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) adminDeleteUser(c echo.Context) error {
	lang := getLanguage(c)

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "http.server",
			Function: "adminDeleteUser",
			Params:   map[string]interface{}{"id": idString},
			Message:  err.Error(),
		})

		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(lang, messages.ParseQueryError),
		}
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		message, code := derrors.HttpError(err)

		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(lang, message),
		}
	}

	return c.NoContent(http.StatusOK)
}
