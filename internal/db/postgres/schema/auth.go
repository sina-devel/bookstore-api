package schema

import "gorm.io/gorm"

type (
	Token struct {
		gorm.Model
		Value  string
		UserID uint
	}
)
