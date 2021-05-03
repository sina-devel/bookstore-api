package models

import "github.com/kianooshaz/bookstore-api/internal/models/types"

type (
	Wallet struct {
		Base
		UserID  uint
		Balance types.Price
		status  types.WalletStatus
	}
)
