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
		Comments      []Comment
		DownloadCount uint
		Pictures      []Picture
		Status        types.BookStatus
		Price         types.Price
	}
)
