package types

type (
	BookStatus uint
)

const (
	_ BookStatus = iota
	BookPending
	BookConfirmed
	BookReject
	BookClose
)
