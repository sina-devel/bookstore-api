package hash

import "testing"

func TestCheckPassword(t *testing.T) {
	password := "secret"
	hashPassword, err := Password(password)
	if err != nil {
		t.Fail()
	}

	if !CheckPassword(password, hashPassword) {
		t.Fail()
	}
}
