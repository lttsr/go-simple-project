package types

type AuthorityType int

const (
	Admin AuthorityType = iota
	User
	Guest
)

func (t AuthorityType) IsAdmin() bool {
	switch t {
	case Admin:
		return true
	default:
		return false
	}
}
func (t AuthorityType) IsGuestGuest() bool {
	switch t {
	case Guest:
		return true
	default:
		return false
	}
}
