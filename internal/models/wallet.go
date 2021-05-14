package models

import "github.com/kianooshaz/bookstore-api/internal/models/types"

type (
	Wallet struct {
		ID      uint               `json:"id"`
		UserID  uint               `json:"user_id"`
		Balance types.Price        `json:"balance"`
		Status  types.WalletStatus `json:"status"`
	}
)
