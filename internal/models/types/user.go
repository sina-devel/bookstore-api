package types

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
)

const (
	_ Gender = iota
	Male
	Female
	More
)

const (
	_ Role = iota
	Basic
	Seller
	Manager
	Admin
)
