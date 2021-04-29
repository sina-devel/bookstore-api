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

	switch strings.ToLower(lang) {
	case "en":
		return EN
	case "fa":
		return FA
	default:
		return EN
	}
}
