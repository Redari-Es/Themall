package controllers

type KouhongController struct {
	BaseController
}

type FendiyeController struct {
	BaseController
}
type YanyingController struct {
	BaseController
}

func (c *KouhongController) Get() {
	c.Data["login"] = "/mall"
	c.Data["about"] = "/mall/about"
	c.Data["jingxuan"] = "/mall"
	c.Data["dacu"] = "/mall"
	c.Data["one"] = "#"
	c.Data["two"] = "/mall/fendiye"
	c.Data["three"] = "/yanying"

	//	c.TplName = "kouhong/kouhong.html"
	c.TplName = "kouhong/kouhong.html"
}

func (c *FendiyeController) Get() {
	c.Data["login"] = "/mall"
	c.Data["about"] = "/mall/about"
	c.Data["jingxuan"] = "/mall"
	c.Data["dacu"] = "/mall"
	c.Data["one"] = "/mall/kouhong"
	c.Data["two"] = "#"
	c.Data["three"] = "/mall/yanying"

	c.TplName = "fendiye/fendiye.html"
}

func (c *YanyingController) Get() {
	c.Data["login"] = "/mall"
	c.Data["about"] = "/mall/about"
	c.Data["jingxuan"] = "/mall"
	c.Data["dacu"] = "/mall"
	c.Data["one"] = "/mall/kouhong"
	c.Data["two"] = "/mall/fendiye"
	c.Data["three"] = "#"

	c.TplName = "yanying/yanying.html"
}
