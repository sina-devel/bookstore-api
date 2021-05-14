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

func TestGetUserByID(t *testing.T) {
	user := newUserTest()

	t.Run("create new record", func(t *testing.T) {
		if err := repoTest.CreateUser(user); err != nil {
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

	t.Run("create new record", func(t *testing.T) {
		if err := repoTest.CreateUser(user); err != nil {
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

	t.Run("create new record", func(t *testing.T) {
		if err := repoTest.CreateUser(user); err != nil {
			t.Fail()
		}
	})

	user.FirstName = "updated"
	user2.ID = uint(rand.Uint32())

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
			want: derrors.New(derrors.KindNotFound, messages.UserNotFound),
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

	t.Run("create new record", func(t *testing.T) {
		if err := repoTest.CreateUser(user); err != nil {
			t.Fail()
		}
	})

	user2.ID = uint(rand.Uint32())

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

	tests := []struct {
		name string
		user *models.User
		want error
	}{
		{
			name: "create user and wallet",
			user: user,
			want: nil,
		},
		{
			name: "username is not unique",
			user: user,
			want: derrors.New(derrors.KindUnexpected, messages.DBError),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repoTest.CreateUser(tt.user)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}

}
