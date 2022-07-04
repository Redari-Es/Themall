package controllers

//"github.com/astaxie/beego"

type ExitController struct {
	//beego.Controller
	AdminController
}
type UserexitController struct {
	//beego.Controller
	ExitController
}

func (c *ExitController) Get() {
	c.DelSession("loginuser")
	c.Data["errmsg"] = "用户已退出，请重新登录o(h_h)o"
	c.Redirect("/mall/login", 302)

}
func (c *UserexitController) Get() {
	c.TplName = "base/exit.html"

}
