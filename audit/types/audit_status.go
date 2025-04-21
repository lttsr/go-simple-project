package types

type AuditStatusType int

const (
	UNPROCESSED AuditStatusType = iota
	PROCESSING
	PROCESSED
	ERROR
)
