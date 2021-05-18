package server

import (
	"github.com/kianooshaz/bookstore-api/pkg/translate"
	"github.com/labstack/echo/v4"
)

func getLanguage(c echo.Context) translate.Language {
	accept := c.Request().Header.Get("Accept-Language")
	switch accept {
	case "fa", "fa-ir", "farsi", "persian":
		return translate.FA
	case "en", "en-us", "english":
		return translate.EN
	default:
		return translate.EN
	}
}
