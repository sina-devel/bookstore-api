package postgres

import (
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"testing"
)

func TestAddUser(t *testing.T) {
	user := newUserTest()

	user2 := newUserTest()

	wallet := &models.Wallet{
		UserID:  0,
		Balance: 0,
		Status:  types.WalletOpen,
	}

	tests := []struct {
		name      string
		user      *models.User
		wallet    *models.Wallet
		wantError bool
	}{
		{
			name:      "create user and wallet",
			user:      user,
			wallet:    wallet,
			wantError: false,
		},
		{
			name:      "username is not unique",
			user:      user,
			wallet:    wallet,
			wantError: true,
		},
		{
			name:      "create user2 and wallet",
			user:      user2,
			wallet:    wallet,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repoTest.AddUser(tt.user, tt.wallet)
			if tt.wantError {
				if err == nil {
					t.Errorf("error got = %v, want %v", err, tt.wantError)
				}
			} else {
				if err != nil {
					t.Errorf("error got = %v, want %v", err, tt.wantError)
				}
			}
		})
	}

}
