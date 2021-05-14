package schema

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
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

func (w *Wallet) ConvertModel() *models.Wallet {
	return &models.Wallet{
		ID:      w.ID,
		UserID:  w.UserID,
		Balance: w.Balance,
		Status:  w.Status,
	}
}

func ConvertWallet(wallet *Wallet) *Wallet {
	return &Wallet{
		Model:   gorm.Model{ID: wallet.ID},
		UserID:  wallet.UserID,
		Balance: wallet.Balance,
		Status:  wallet.Status,
	}
}
