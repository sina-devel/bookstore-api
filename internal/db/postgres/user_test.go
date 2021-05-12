package postgres

import (
	"errors"
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/random"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
	"math/rand"
	"testing"
)

func newUserTest() *models.User {
	return &models.User{
		Username:    random.String(8),
		FirstName:   random.String(8),
		LastName:    random.String(8),
		Email:       random.String(5) + "@" + random.String(3) + "." + random.String(3),
		PhoneNumber: "0912" + random.StringWithCharset(7, "0123456789"),
		Gender:      types.Male,
		Role:        types.Basic,
	}
}

func TestGetUserByID(t *testing.T) {
	user := newUserTest()

	t.Run("insert new record", func(t *testing.T) {
		if rows := repoTest.db.Create(user).
			RowsAffected; rows != 1 {
			t.Fail()
		}
	})

	test := []struct {
		name string
		id   uint
		want error
	}{
		{
			name: "get user by id",
			id:   user.ID,
			want: nil,
		},
		{
			name: "user not found",
			id:   uint(rand.Uint64()),
			want: derrors.New(derrors.KindNotFound, messages.UserNotFound),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repoTest.GetUserByID(tt.id)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}
}

func TestGetUserByUsername(t *testing.T) {
	user := newUserTest()

	t.Run("insert new record", func(t *testing.T) {
		if rows := repoTest.db.Create(user).
			RowsAffected; rows != 1 {
			t.Fail()
		}
	})

	tests := []struct {
		name     string
		username string
		want     error
	}{
		{
			name:     "get user by username",
			username: user.Username,
			want:     nil,
		},
		{
			name:     "user not found",
			username: random.String(10),
			want:     derrors.New(derrors.KindNotFound, messages.UserNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repoTest.GetUserByUsername(tt.username)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	user := newUserTest()
	user2 := newUserTest()

	t.Run("insert new record", func(t *testing.T) {
		if rows := repoTest.db.Create(user).
			RowsAffected; rows != 1 {
			t.Fail()
		}
	})

	user.FirstName = random.String(5)
	user2.ID = uint(rand.Intn(5))

	tests := []struct {
		name string
		user *models.User
		want error
	}{
		{
			name: "update user",
			user: user,
			want: nil,
		},
		{
			name: "user not found",
			user: user2,
			want: derrors.New(derrors.KindUnexpected, messages.UserNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repoTest.UpdateUser(tt.user)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}
}

func TestDeleteUserByID(t *testing.T) {
	user := newUserTest()
	user2 := newUserTest()

	t.Run("insert new record", func(t *testing.T) {
		if rows := repoTest.db.Create(user).
			RowsAffected; rows != 1 {
			t.Fail()
		}
	})

	user2.ID = uint(rand.Intn(4))

	tests := []struct {
		name string
		user *models.User
		want error
	}{
		{
			name: "delete user",
			user: user,
			want: nil,
		},
		{
			name: "user not found",
			user: user2,
			want: derrors.New(derrors.KindNotFound, messages.UserNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repoTest.DeleteUser(tt.user)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}
}

func TestAddUser(t *testing.T) {

	user := newUserTest()

	user2 := newUserTest()

	wallet := &models.Wallet{
		UserID:  0,
		Balance: 0,
		Status:  types.WalletOpen,
	}

	tests := []struct {
		name   string
		user   *models.User
		wallet *models.Wallet
		want   error
	}{
		{
			name:   "create user and wallet",
			user:   user,
			wallet: wallet,
			want:   nil,
		},
		{
			name:   "username is not unique",
			user:   user,
			wallet: wallet,
			want:   derrors.New(derrors.KindUnexpected, messages.DBError),
		},
		{
			name:   "create user2 and wallet",
			user:   user2,
			wallet: wallet,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repoTest.AddUser(tt.user, tt.wallet)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}

}
