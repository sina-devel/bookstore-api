package models

type (
	Comment struct {
		ID          uint
		UserID      uint
		Text        string
		BookID      uint
		FullName    string
		IsConfirmed bool
	}
)
