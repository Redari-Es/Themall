package controllers

type ContactController struct {
	BaseController
}

func (c *ContactController) Get() {
	c.Data["login"] = "/mall"
	c.Data["about"] = "/mall/about"
	c.Data["jingxuan"] = "/mall"
	c.Data["dacu"] = "/mall"
	c.Data["one"] = "/mall"
	c.Data["two"] = "/mall"
	c.Data["three"] = "/mall"

	c.TplName = "contact/contact.html"
}
