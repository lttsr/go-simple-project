package audit

import (
	"beego-app/audit/types"
	"time"

	"github.com/beego/beego/v2/client/orm"
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

func Register(rep orm.Ormer, param AuditActorRegsister) (*AuditActor, error) {
	audit := &AuditActor{
		ActorId:         0,
		Category:        param.Category,
		Message:         param.Message,
		ErrorReason:     param.ErrorReason,
		AuditStatusType: param.AuditStatusType,
	}

	_, err := rep.Insert(audit)
	if err != nil {
		return nil, err
	}

	return audit, nil
}
