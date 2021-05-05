package types

import (
	"fmt"
	"github.com/kianooshaz/bookstore-api/pkg/errors"
	"github.com/kianooshaz/bookstore-api/pkg/translator/messages"
)

type (
	UserStatus uint

	Gender uint

	Role uint
)

const (
	_ UserStatus = iota
	UserActive
	UserInactive
	UserBan

	_ Gender = iota
	Male
	Female
	More

	_ Role = iota
	Basic
	Seller
	Manager
	Admin
)

var (
	userStatus = map[UserStatus]string{
		UserActive:   "Active",
		UserInactive: "Inactive",
		UserBan:      "Ban",
	}

	gender = map[Gender]string{
		Male:   "Male",
		Female: "Female",
		More:   "More",
	}

	role = map[Role]string{
		Basic:   "Basic",
		Seller:  "Seller",
		Manager: "Manager",
		Admin:   "Admin",
	}
)

func (u UserStatus) String() string {
	if s, ok := userStatus[u]; ok {
		return s
	}

	return fmt.Sprintf("UserStatus(%d)", u)
}

func (u UserStatus) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}

func (u *UserStatus) UnmarshalText(b []byte) error {
	for i, v := range userStatus {
		if v == string(b) {
			*u = i
			return nil
		}
	}

	return errors.New(errors.KindInvalid, messages.InvalidUserStatus)
}

func (g Gender) String() string {
	if s, ok := gender[g]; ok {
		return s
	}

	return fmt.Sprintf("Gender(%d)", g)
}

func (g Gender) MarshalText() ([]byte, error) {
	return []byte(g.String()), nil
}

func (g *Gender) UnmarshalText(b []byte) error {
	for i, v := range gender {
		if v == string(b) {
			*g = i
			return nil
		}
	}

	return errors.New(errors.KindInvalid, messages.InvalidGender)
}

func (r Role) String() string {
	if s, ok := role[r]; ok {
		return s
	}

	return fmt.Sprintf("Role(%d)", r)
}

func (r Role) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Role) UnmarshalText(b []byte) error {
	for i, v := range role {
		if v == string(b) {
			*r = i
			return nil
		}
	}

	return errors.New(errors.KindInvalid, messages.InvalidRole)
}
