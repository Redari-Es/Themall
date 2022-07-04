package controllers

type AboutController struct {
	BaseController
}

func (c *AboutController) Get() {
	// TODO: navbar  <24-12-21, Redari-Es> //
	if c.IsLogin {
		c.Data["login"] = c.Loginuser

	} else {
		c.Data["login"] = "/mall/login"
	}
	c.Data["about"] = "about"
	c.Data["love"] = "/mall"
	c.Data["jingxuan"] = "/mall"
	c.Data["dacu"] = "/mall"
	c.Data["one"] = "/mall"
	c.Data["two"] = "/mall"
	c.Data["three"] = "/mall"

	c.TplName = "aboutus/aboutus.html"
}
