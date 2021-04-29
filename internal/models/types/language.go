package types

type (
	Language string
)

const (
	FA Language = "fa"
	EN Language = "en"
)

func GetLanguage(lang string) Language {
	switch lang {
	case "en":
		return EN
	case "fa":
		return FA
	default:
		return EN
	}
}
