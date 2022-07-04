package controllers

import "fmt"

type IsloginController struct {
	BaseController
}

type IsmallController struct {
	BaseController
}

func (c *IsloginController) Get() {

	c.Data["username"] = "helo"
	c.Data["passwd"] = "helo"
	c.Data["email"] = "helo"
	c.TplName = "login.html"
}
func (c *IsmallController) Get() {
	username := c.Ctx.Input.Param(":user")
	fmt.Println("This is test now!:", username)
	if username == c.GetSession("loginuser") {
		c.Data["login"] = "mall/" + username
		c.Data["about"] = "mall/about"
		c.Data["jingxuan"] = "#jingxuan"
		c.Data["dacu"] = "#dacu"
		c.Data["one"] = "#one"
		c.Data["two"] = "#two"
		c.Data["three"] = "#three"

		c.TplName = "mall/mall.html"

	} else {
		// TODO: navbar  <24-12-21, Redari-Es> //

		c.TplName = "login/login.html"

	}

}
