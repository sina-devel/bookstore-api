package models

type (
	Category struct {
		Base
		Name        string
		Description string
		Slug        string
	}
)
