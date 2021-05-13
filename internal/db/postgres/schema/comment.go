package schema

import "gorm.io/gorm"

type (
	Comment struct {
		gorm.Model
		UserID      uint
		Text        string
		BookID      uint
		FullName    string
		IsConfirmed bool
	}
)
