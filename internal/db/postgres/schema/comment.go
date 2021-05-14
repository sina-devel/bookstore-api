package schema

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"gorm.io/gorm"
)

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

func (c *Comment) ConvertModel() *models.Comment {
	return &models.Comment{
		ID:          c.ID,
		UserID:      c.UserID,
		Text:        c.Text,
		BookID:      c.BookID,
		FullName:    c.FullName,
		IsConfirmed: c.IsConfirmed,
	}
}

func ConvertComment(comment *models.Comment) *Comment {
	return &Comment{
		Model:       gorm.Model{ID: comment.ID},
		UserID:      comment.UserID,
		Text:        comment.Text,
		BookID:      comment.BookID,
		FullName:    comment.FullName,
		IsConfirmed: comment.IsConfirmed,
	}
}
