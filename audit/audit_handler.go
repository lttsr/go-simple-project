package audit

import (
	"beego-app/audit/types"
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/sirupsen/logrus"
)

var LoggerActor = logrus.New()
var LoggerEvent = logrus.New()

func Init() {
	LoggerActor.SetOutput(os.Stdout)
	LoggerActor.SetLevel(logrus.DebugLevel)
	LoggerActor.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	LoggerEvent.SetOutput(os.Stdout)
	LoggerEvent.SetLevel(logrus.DebugLevel)
	LoggerEvent.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
}

func Audit(message string, categroy string, fn func() (interface{}, error)) (interface{}, error) {
	var margemessage = categroy + "." + message
	LoggerActor.Info("[開始]" + margemessage)
	result, err := fn()
	//systemuser registeraudit
	param := AuditActorRegsister{
		// ActorId:         loginUsrId,
		Category:        categroy,
		Message:         message,
		ErrorReason:     "",
		AuditStatusType: types.AuditStatusType(0),
	}
	if err != nil {
		param.ErrorReason = err.Error()
		param.AuditStatusType = types.AuditStatusType(4)
		LoggerActor.Error("[例外]" + margemessage)
		return result, err
	}
	LoggerActor.Info("[正常終了]" + margemessage)
	param.AuditStatusType = types.AuditStatusType(3)
	Register(orm.NewOrm(), param)
	return result, err
}

type AuditActorRegsister struct {
	Category        string
	Message         string
	ErrorReason     string
	AuditStatusType types.AuditStatusType
}
