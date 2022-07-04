package controllers

import "fmt"

type MallController struct {
	//beego.Controller
	BaseController
}

func (c *MallController) Get() {

	if c.IsLogin {
		fmt.Println("用户正在访问中：", c.Loginuser)
		fmt.Println("登录状态：", c.IsLogin)

	} else {
		fmt.Println("尚未有用户登录：", c.Loginuser)
		fmt.Println("登录状态：", c.IsLogin)

	}
	// TODO: navbar  <28-12-21, Redari-Es> //

	c.TplName = "mall/home.html"
}
