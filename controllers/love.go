package controllers

type LoveController struct {
	BaseController
}

func (c *LoveController) Get() {
	// TODO: navbar  <24-12-21, Redari-Es> //
	c.Data["login"] = "/mall"
	c.Data["jingxuan"] = "/mall"
	c.Data["dacu"] = "/mall"
	c.Data["one"] = "/mall"
	c.Data["two"] = "/mall"
	c.Data["three"] = "/mall"
	c.Data["love"] = "#"

	c.TplName = "love/love.html"
}
