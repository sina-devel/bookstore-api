package schema

import "gorm.io/gorm"

type (
	Picture struct {
		gorm.Model
		Name   string
		Alt    string
		BookID uint
	}
)
