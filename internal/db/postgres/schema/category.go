package schema

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"gorm.io/gorm"
)

type (
	Category struct {
		gorm.Model
		Name        string
		Description string
		Slug        string
	}
)

func (c *Category) ConvertModel() *models.Category {
	return &models.Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Slug:        c.Slug,
	}
}

func ConvertCategory(category *models.Category) *Category {
	return &Category{
		Model:       gorm.Model{ID: category.ID},
		Name:        category.Name,
		Description: category.Description,
		Slug:        category.Slug,
	}
}
