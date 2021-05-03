package models

type (
	Comment struct {
		Base
		UserID      uint
		Text        string
		BookID      uint
		FullName    string
		IsConfirmed bool
	}
)
