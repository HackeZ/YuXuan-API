package controllers

import (
	"YuXuanAPI/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations for YuXuanAdmin
type YxAdminController struct {
	beego.Controller
}

func (c *YxAdminController) URLMapping() {
	c.Mapping("Post", c.Post)
	// c.Mapping("GetOne", c.GetOne)
	// c.Mapping("GetAll", c.GetAll)
	// c.Mapping("Put", c.Put)
	// c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create YxAdmin
// @Param	body		body 	models.YxAdmin	true		"body for YxAdmin content"
// @Success 201 {int} models.YxAdmin
// @Failure 403 body is empty
// @router / [post]
func (c *YxAdminController) POST() {
	var v models.YxAdmin
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddYxAdmin(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
