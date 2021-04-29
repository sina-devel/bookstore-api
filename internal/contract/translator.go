package contract

import "github.com/kianooshaz/bookstore-api/internal/models/types"

type (
	Translator interface {
		Translate(lang types.Language, key string) string
		TranslateEn(key string) string
	}
)
