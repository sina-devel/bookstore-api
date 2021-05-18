package hash

import (
	"errors"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
	"golang.org/x/crypto/bcrypt"
)

func Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return derrors.New(derrors.KindInvalid, messages.UsernameOrPasswordIsIncorrect)
		}

		return derrors.New(derrors.KindUnexpected, messages.GeneralError)
	}

	return nil
}
