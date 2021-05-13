package schema

import (
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"gorm.io/gorm"
)

type (
	Wallet struct {
		gorm.Model
		UserID  uint
		Balance types.Price
		Status  types.WalletStatus
	}
)
