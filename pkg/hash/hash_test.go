package hash

import (
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/random"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	password := random.String(6)
	hashPassword, err := Password(password)
	if err != nil {
		t.Fail()
	}

	test := []struct {
		name     string
		password string
		hash     string
		want     error
	}{
		{
			name:     "correct test case",
			password: password,
			hash:     hashPassword,
			want:     nil,
		},
		{
			name:     "incorrect test case",
			password: random.String(7),
			hash:     hashPassword,
			want:     derrors.New(derrors.KindInvalid, messages.UsernameOrPasswordIsIncorrect),
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPassword(tt.password, tt.hash)
			if err != tt.want {
				t.Fail()
			}
		})
	}
}
