package models

import "github.com/kianooshaz/bookstore-api/internal/models/types"

type (
	Wallet struct {
		ID      uint
		UserID  uint
		Balance types.Price
		Status  types.WalletStatus
	}
)
