package routers

import (
	"beego/BeeRobot/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/cq_event/", &controllers.CQController{})
	beego.Router("/cq_event/list/", &controllers.CQController{}, "*:ListCQreq")
	beego.Router("/cq_event/content/", &controllers.CQController{}, "*:ContentCQreq")
}
