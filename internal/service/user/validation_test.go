package user

import (
	"errors"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/random"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
	"testing"
)

func TestValidateUsername(t *testing.T) {
	setupTest(t)
	defer tearDownTest()

	test := []struct {
		name     string
		username string
		want     error
	}{
		{
			name:     "correct username",
			username: "admin",
			want:     nil,
		},
		{
			name:     "invalid username for min",
			username: random.String(2),
			want:     derrors.New(derrors.KindInvalid, messages.InvalidUsernameLength),
		},
		{
			name:     "invalid username for max",
			username: random.String(55),
			want:     derrors.New(derrors.KindInvalid, messages.InvalidUsernameLength),
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := serviceTest.validateUsername(tt.username)
			if !errors.Is(tt.want, err) {
				t.Error()
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	setupTest(t)
	defer tearDownTest()

	test := []struct {
		name     string
		password string
		want     error
	}{
		{
			name:     "correct password",
			password: "Golang&*526sw",
			want:     nil,
		},
		{
			name:     "invalid password for upper letter",
			password: "sdwdm%&52",
			want:     derrors.New(derrors.KindInvalid, messages.InvalidPassword),
		},
		{
			name:     "invalid password for symbol",
			password: "CdaVDlow626",
			want:     derrors.New(derrors.KindInvalid, messages.InvalidPassword),
		},
		{
			name:     "invalid password for letter",
			password: "15114&%@621848",
			want:     derrors.New(derrors.KindInvalid, messages.InvalidPassword),
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := serviceTest.validatePassword(tt.password)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}
}
