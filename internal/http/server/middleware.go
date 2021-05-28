package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
	"github.com/labstack/echo/v4"
	"net/http"
)

func middlewarePermission(h *handler, roles ...types.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			lang := getLanguage(c)

			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				h.logger.Error(&log.Field{
					Section:  "server.middleware",
					Function: "middlewarePermission",
					Message:  h.translator.TranslateEn(messages.InvalidToken),
				})

				return &echo.HTTPError{
					Code:    http.StatusUnauthorized,
					Message: h.translator.Translate(lang, messages.InvalidToken),
				}
			}

			user, ok := token.Claims.(*models.Claims)
			if !ok {
				h.logger.Error(&log.Field{
					Section:  "server.middleware",
					Function: "middlewarePermission",
					Message:  h.translator.TranslateEn(messages.InvalidToken),
				})

				return &echo.HTTPError{
					Code:    http.StatusUnauthorized,
					Message: h.translator.Translate(lang, messages.InvalidToken),
				}
			}

			for _, role := range roles {
				if user.Role == role {
					return next(c)
				}
			}

			h.logger.Error(&log.Field{
				Section:  "server.middleware",
				Function: "middlewarePermission",
				Message:  h.translator.TranslateEn(messages.NotPermission),
			})

			return &echo.HTTPError{
				Code:    http.StatusForbidden,
				Message: h.translator.Translate(lang, messages.NotPermission),
			}
		}
	}
}
