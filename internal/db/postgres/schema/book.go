package schema

import (
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"gorm.io/gorm"
)

type (
	Book struct {
		gorm.Model
		Name          string
		Description   string
		File          string
		SellerID      uint
		CategoryID    uint
		Comments      []Comment `gorm:"foreignKey:BookID"`
		DownloadCount uint
		Pictures      []Picture `gorm:"foreignKey:BookID"`
		Status        types.BookStatus
		Price         types.Price
	}
)