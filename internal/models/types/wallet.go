package types

import (
	"fmt"
	"github.com/kianooshaz/bookstore-api/pkg/errors"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

type (
	WalletStatus uint
)

const (
	_ WalletStatus = iota
	WalletOpen
	WalletClose
	WalletInactive
)

var (
	walletStatus = map[WalletStatus]string{
		WalletOpen:     "open",
		WalletClose:    "close",
		WalletInactive: "inactive",
	}
)

func (w WalletStatus) String() string {
	if s, ok := walletStatus[w]; ok {
		return s
	}

	return fmt.Sprintf("WalletStatus(%d)", w)
}

func (w WalletStatus) MarshalText() ([]byte, error) {
	return []byte(w.String()), nil
}

func (w *WalletStatus) UnmarshalText(b []byte) error {
	for i, v := range walletStatus {
		if v == string(b) {
			*w = i
			return nil
		}
	}

	return errors.New(errors.KindInvalid, messages.InvalidWalletStatus)
}
