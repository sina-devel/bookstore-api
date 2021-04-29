package translator

import "strings"

type (
	Translator interface {
		Translate(lang Language, key string) string
		TranslateEn(key string) string
	}

	Language string
)

const (
	EN Language = "en"
	FA Language = "fa"
)

func GetLanguage(lang string) Language {
	switch strings.ToLower(lang) {
	case "en":
		return EN
	case "fa":
		return FA
	default:
		return EN
	}
}
