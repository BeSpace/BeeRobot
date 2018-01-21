package controllers

import (
	"beego/BeeRobot/models"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
)

type CQController struct {
	beego.Controller
}

/* 这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类。*/
func (this *CQController) Prepare() {

}

func (this *CQController) Post() {

	this.EnableRender = false // 接口，统一关闭模版渲染

	requestBody := string(this.Ctx.Input.RequestBody)
	receivedSHA1 := string(string(this.Ctx.Input.Header("x-signature")))

	isRaw := checkSHA1(requestBody, receivedSHA1)

	if isRaw {
		CQreq := models.CQreq{}
		CQreq.Message = string(requestBody)
		models.SaveCQreq(&CQreq)
	}

}

func (this *CQController) ListCQreq() {

	var params models.BaseQueryParam

	params.Page, _ = this.GetInt("page")
	params.Limit, _ = this.GetInt("limit")

	data, total := models.PageListCQreg(&params)
	result := make(map[string]interface{})
	result["total"] = total
	result["data"] = data
	result["status"] = 200
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *CQController) ContentCQreq() {
	this.EnableRender = true
	this.TplName = "content/QQrobot/reqmessage.tpl"
	this.LayoutSections = make(map[string]string)
	this.Layout = "index.tpl"
	this.LayoutSections["Header"] = "common/header.tpl"
	this.LayoutSections["SiderBar"] = "sider/sider.tpl"
	this.LayoutSections["FootBar"] = "footer/footer.tpl"
}

/**
 * [checkSHA1]
 * @param  {[type]} requestBody  string        [description]
 * @param  {[type]} receivedSHA1 string)       (isRaw        bool [description]
 * @return {[type]}              [description]
 */
func checkSHA1(requestBody string, receivedSHA1 string) (isRaw bool) {
	secret := beego.AppConfig.String("secret")

	//hmac ,use sha1
	key := []byte(secret)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(requestBody))

	// Validate signature.
	isRaw = (string(fmt.Sprintf("sha1=%x", mac.Sum(nil))) == receivedSHA1)
	return
}
