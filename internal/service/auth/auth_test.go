package auth

import (
	"errors"
	"github.com/kianooshaz/bookstore-api/internal/models"
	"testing"
)

func TestGenerateAccessToken(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()

	tests := []struct {
		name string
		user *models.User
		want error
	}{
		{
			name: "correct test",
			user: user,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := serviceTest.GenerateAccessToken(tt.user)
			t.Log("access_token: " + token)
			if !errors.Is(err, tt.want) {
				t.Fail()
			}
		})
	}
}

func TestGenerateRefreshToken(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()

	tests := []struct {
		name string
		user *models.User
		want error
	}{
		{
			name: "correct test",
			user: user,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := serviceTest.GenerateRefreshToken(tt.user)
			t.Log("refresh_token: " + token)
			if !errors.Is(err, tt.want) {
				t.Fail()
			}
		})
	}
}
