package models

type (
	Picture struct {
		Base
		Name   string
		Alt    string
		BookID uint
	}
)
