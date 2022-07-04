package controllers

import (
	"fmt"
	"mall/models/utils"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	Loginuser interface{}
}

func (c *BaseController) Prepare() {
	loginuser := c.GetSession("loginuser")
	fmt.Println("开始访问：", loginuser)
	fendiye := "#"
	kouhong := "#"
	yanying := "#"
	love := "mall/login"
	login := "mall/login"

	if loginuser != nil {
		c.IsLogin = true
		c.Loginuser = loginuser
		uname := utils.Strval(loginuser)
		c.Data["username"] = uname
		login = "mall/" + uname
		love = "mall/love"
		//商品内容
		fendiye = "mall/fendiye"
		kouhong = "mall/kouhong"
		yanying = "mall/yanying"

		//c.TplName = "mall/home.html"

	} else {
		c.IsLogin = false
		fmt.Println("用户尚未登录！")
		//	c.Redirect("mall/login", 302)

		//c.TplName = "mall/home.html"
		//		c.Redirect("/mall/login", 302)
	}
	// TODO: navbar  <28-12-21, Redari-Es> //
	c.Data["login"] = login
	c.Data["about"] = "mall/about"
	c.Data["love"] = love
	c.Data["contact"] = "mall/contact"
	c.Data["jingxuan"] = "#jingxuan"
	c.Data["dacu"] = "#dacu"
	c.Data["one"] = "#one"
	c.Data["two"] = "#two"
	c.Data["three"] = "#three"
	//商品内容
	c.Data["fendiye"] = fendiye
	c.Data["kouhong"] = kouhong
	c.Data["yanying"] = yanying

	c.Data["Islogin"] = c.IsLogin
	//c.Redirect("/mall/:loginuser", 200)

}
