package routers

import (
	"mall/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//	beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.BaseController{})
	//mall
	beego.Router("/mall", &controllers.MallController{})
	beego.Router("/mall/test", &controllers.IsloginController{})
	beego.Router("/mall/:user", &controllers.AdminController{})
	//退出
	beego.Router("/mall/:user/exit", &controllers.UserexitController{})
	beego.Router("/mall/exit", &controllers.ExitController{})
	//beego.Router("/mall/:user", &controllers.IsmallController{})

	//登录注册
	beego.Router("/mall/login", &controllers.UserController{})
	beego.Router("/mall/signin", &controllers.SigninController{})
	beego.Router("/mall/signup", &controllers.SignupController{})
	//订单

	beego.Router("/mall/love", &controllers.LoveController{})
	//联系
	beego.Router("/mall/contact", &controllers.ContactController{})
	//关于我们
	beego.Router("/mall/about", &controllers.AboutController{})
	//商品
	beego.Router("/mall/yanying", &controllers.YanyingController{})
	beego.Router("/mall/fendiye", &controllers.FendiyeController{})
	beego.Router("/mall/kouhong", &controllers.KouhongController{})
}
