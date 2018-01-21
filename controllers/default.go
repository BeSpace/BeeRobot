package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "content/QQrobot/reqmessage.tpl"
	this.LayoutSections = make(map[string]string)
	this.Layout = "index.tpl"
	this.LayoutSections["Header"] = "common/header.tpl"
	this.LayoutSections["SiderBar"] = "sider/sider.tpl"
	this.LayoutSections["FootBar"] = "footer/footer.tpl"
}
