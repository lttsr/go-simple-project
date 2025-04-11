package types

type StatusType int

const (
	Temporary AuthorityType = iota
	Valid
	Invalid
)

func (t AuthorityType) IsValid() bool {
	switch t {
	case Valid:
		return true
	default:
		return false
	}
}
