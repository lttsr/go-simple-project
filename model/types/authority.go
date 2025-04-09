package types

type AuthorityType int

const (
	Admin AuthorityType = iota
	User
	Guest
)

func (a AuthorityType) String() string {
	return [...]string{"Admin", "User", "Guest"}[a]
}
