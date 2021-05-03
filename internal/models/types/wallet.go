package types

type (
	WalletStatus uint
)

const (
	_ WalletStatus = iota
	WalletOpen
	WalletClose
	WalletInactive
)
