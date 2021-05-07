package types

import (
	"fmt"
	"github.com/kianooshaz/bookstore-api/pkg/errors"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

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

var (
	bookStatus = map[BookStatus]string{
		BookPending:   "pending",
		BookConfirmed: "confirmed",
		BookReject:    "reject",
		BookClose:     "close",
	}
)

func (b BookStatus) String() string {
	if s, ok := bookStatus[b]; ok {
		return s
	}

	return fmt.Sprintf("BookStatus(%d)", b)
}

func (b BookStatus) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *BookStatus) UnmarshalText(by []byte) error {
	for i, v := range bookStatus {
		if v == string(by) {
			*b = i
			return nil
		}
	}

	return errors.New(errors.KindInvalid, messages.InvalidBookStatus)
}
