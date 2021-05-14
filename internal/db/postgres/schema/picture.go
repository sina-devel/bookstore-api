package schema

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"gorm.io/gorm"
)

type (
	Picture struct {
		gorm.Model
		Name   string
		Alt    string
		BookID uint
	}
)

func (p *Picture) ConvertModel() *models.Picture {
	return &models.Picture{
		ID:     p.ID,
		Name:   p.Name,
		Alt:    p.Alt,
		BookID: p.BookID,
	}
}

func ConvertPicture(picture *models.Picture) *Picture {
	return &Picture{
		Model:  gorm.Model{ID: picture.ID},
		Name:   picture.Name,
		Alt:    picture.Alt,
		BookID: picture.BookID,
	}
}
