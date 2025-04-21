package audit

import (
	"beego-app/audit/types"
	"time"
)

type AuditActor struct {
	AuditId         int                   `orm:"auto"`
	ActorId         int                   `orm:"size(30)"`
	Category        string                `orm:"size(30)"`
	Message         string                `orm:"size(100)"`
	AuditStatusType types.AuditStatusType `orm:"default(0)"`
	ErrorReason     string                `orm:"size(100)"`
	RegisterId      string                `orm:"auto_now_add"`
	RegisterDate    time.Time             `orm:"auto_now_add"`
	UpdateId        string                `orm:"auto_now_add"`
	UpdateDate      time.Time             `orm:"auto_now_add"`
}
