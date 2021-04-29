package types

import "strings"

type (
	Language string
)

const (
	FA Language = "fa"
	EN Language = "en"
)

func GetLanguage(lang string) Language {

	lang = strings.ToLower(lang)

	switch lang {
	case "en":
		return EN
	case "fa":
		return FA
	default:
		return EN
	}
}
