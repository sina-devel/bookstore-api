package server

import (
	"github.com/kianooshaz/bookstore-api/pkg/translate"
	"github.com/labstack/echo/v4"
	"sort"
	"strconv"
	"strings"
)

type langQ struct {
	lang string
	q    float64
}

func getLanguage(c echo.Context) translate.Language {
	acceptLanguages := c.Request().Header.Get("Accept-Language")

	lqs := parseAcceptLanguage(acceptLanguages)

	for _, lq := range lqs {
		switch lq.lang {
		case "fa", "fa-ir":
			return translate.FA
		case "en", "en-us", "en-gb":
			return translate.EN
		}
	}

	return translate.EN
}

func parseAcceptLanguage(acceptLanguages string) []langQ {
	var lqs []langQ

	languages := strings.Split(acceptLanguages, ",")
	for _, language := range languages {
		language = strings.Trim(language, " ")

		langWithQ := strings.Split(language, ";")
		if len(langWithQ) == 1 {
			lq := langQ{langWithQ[0], 1}
			lqs = append(lqs, lq)
		} else {
			valueQ := strings.Split(langWithQ[1], "=")
			q, err := strconv.ParseFloat(valueQ[1], 64)
			if err != nil {
				continue
			}
			lq := langQ{langWithQ[0], q}
			lqs = append(lqs, lq)
		}
	}

	sort.SliceStable(lqs, func(i, j int) bool {
		return lqs[i].q > lqs[j].q
	})

	return lqs
}
