package hash

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsErrorInvalidPassword(err error) bool {
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return true
	}

	return false
}
