package controller

import (
	"beego-app/audit"
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type DefaultController struct {
	web.Controller
}

// Bind the request values to struct.
func (b *DefaultController) BindRequest(req interface{}) error {
	body := b.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, req); err != nil {
		audit.Log.Error("Failed to bind request:", err)
	}
	return nil
}

func (b *DefaultController) BadRequest(message string) {
	b.Ctx.Output.SetStatus(400)
	b.Data["json"] = map[string]string{"error": message}
	b.ServeJSON()
}

func (b *DefaultController) NotFound(message string) {
	b.Ctx.Output.SetStatus(404)
	b.Data["json"] = map[string]string{"error": message}
	b.ServeJSON()
}
