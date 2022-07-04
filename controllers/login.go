package controllers

import (
	"fmt"
	"mall/models"
	"mall/models/utils"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

type SigninController struct {
	//BaseController
	//beego.Controller
	UserController
}

type SignupController struct {
	beego.Controller
}

var v = utils.ValiCode(4)

var msg string
var msg2 string

// TODO: login页面  <22-12-21, Redari-Es> //
func (c *UserController) Get() {
	c.Data["errmsg"] = msg
	c.Data["vali"] = v
	//c.Data["errmsg"] = "Plz SignIn"
	c.TplName = "login/login.html"
}

// TODO: form  <22-12-21, Redari-Es> //
type user struct {
	Id       int         `form:"-"`
	Name     interface{} `form:"username"`
	Passwd   string      `form:"password"`
	RePasswd string      `form:"repassword"`
	Email    string      `form:"email"`          // Email 字段需要符合邮箱格式，并且最大长度不能大于 30 个字符
	Vali     string      `form:"vali"`           // Email 字段需要符合邮箱格式，并且最大长度不能大于 30 个字符
	Age      int         `valid:"Range(1, 140)"` // 1 <= Age <= 140，超出此范围即为不合法
	Mobile   string      `valid:"mobile"`        // Mobile 必须为正确的手机号
}

// TODO: login  <22-12-21, Redari-Es> //
/*
func (c *UserController) Post() {
	u := user{}
	if err := c.ParseForm(&u); err == nil {
		//handle erro
	}
	//	helo := c.GetString("username")
	//	if helo == "" {
	//		c.Ctx.WriteString("hahaha")
	//		return
	//	}
	c.Data["username"] = u.Name
	if u.RePasswd != u.Passwd {
		c.Ctx.WriteString("confirm the password!!!")

	}
	c.Data["passwd"] = u.Passwd
	c.Data["repasswd"] = u.RePasswd
	c.Data["email"] = u.Email

	c.TplName = "user.html"
}
*/
// TODO: 数据库操作  <29-12-21, Redari-Es> //

// TODO: 登录表单  <22-12-21, Redari-Es> //
func (c *SigninController) Post() {
	//表单
	u := user{}
	o := orm.NewOrm()
	qs := new(models.User)
	//c.Data["vali"] = v
	if err := c.ParseForm(&u); err != nil {
		//c.Data["errmsg"] = "不能为空"
		msg = "不能为空"
		fmt.Println("登录：0")
		fmt.Println("登录：表单为空")
		fmt.Println("登录：00")
		//handle erro
	} else if u.Name == nil {
		//c.Data["errmsg"] = "请输入用户名"
		msg = "请输入用户名"
		fmt.Println("登录：01->用户名为空")
	} else if u.Passwd == "" {
		//c.Data["errmsg"] = "请输入密码"
		msg = "请输入密码"
		fmt.Println("登录：02->密码为空")
	} else if u.Vali != "" {
		vali := strings.ToLower(u.Vali)
		if vali != strings.ToLower(utils.Strval(v)) {
			//c.Data["errmsg"] = "验证码错误"
			msg = "验证码错误"
			fmt.Println("登录：03-> 验证码错误")
			//c.Redirect("/mall/login", 302)
			/*} else if vali == "" {
			c.Data["errmsg"] = "请输入验证码"
			fmt.Println("登录：05-> 验证码为空")
			*/
		} else {
			uname := utils.Trim(utils.Strval(u.Name))
			qs.Name = uname
			pwd := utils.MD5(strings.TrimSpace(u.Passwd))
			//			qs.Passwd = pwd
			//cx := o.QueryTable("User")

			if err := o.Read(qs, "Name"); err != nil {
				msg = "未有该用户，请注册"
				fmt.Println(("登录：011->未有该用户"))

			} else {
				if qs.Passwd != pwd {
					msg = "密码错误，请重新输入"
					fmt.Println("登录：022->密码错误")
				} else {
					c.Data["username"] = u.Name
					c.Data["passwd"] = u.Passwd
					c.SetSession("loginuser", u.Name)
					c.Data["errmsg"] = "Let's Go Go Go !"
					msg = "用户已登录"
					fmt.Println("--------------------START------------------")
					fmt.Println("登录用户为:", c.GetSession("loginuser"))
					fmt.Println("验证码正确!!!")
					fmt.Println("系统验证码为：", v)
					fmt.Println("用户验证码为：", u.Vali)
					fmt.Println("登录：04-> 登录成功")
					fmt.Println("跳转mall")
					fmt.Println("--------------------END--------------------")
					//if success Redirect to the mall
					c.Redirect("/mall", 302)

				}

			}
		}

		// TODO:   <28-12-21, Redari-Es> //
		//fmt.Println("登录：1验证码为空")
		//c.Data["errmsg"] = "请输入验证码哦!"

		//c.TplName = "login/login.html"

	} else {

		fmt.Println("登录：06-> 其他错误")
		//c.Data["errmsg"] = "验证码不能为空"
		msg = "验证码不能为空"
		c.Redirect("/mall/login", 303)

	}
	fmt.Println("登录：07")
	c.Redirect("/mall/login", 303)
	//c.TplName = "login/login.html"
	//	c.Redirect("mall/test", 200)
	//	c.Ctx.Redirect(303, "/")

	//	helo := c.GetString("username")
	//	if helo == "" {
	//		c.Ctx.WriteString("hahaha")
	//		return
	//	}

}

func (c *SignupController) Post() {
	u := user{}
	o := orm.NewOrm()
	qs := new(models.User)
	// TODO: 格式化  <28-12-21, Redari-Es> //
	if err := c.ParseForm(&u); err != nil {
		msg2 = "请输入内容 (Ovo)"
		fmt.Println("注册：00")
		//handle erro
	} else if u.Name == nil {
		msg2 = "请输入用户名(ovO)"
		fmt.Println("注册：01")
	} else if u.Email == "" {
		msg2 = "请输入邮箱"
		fmt.Println("注册：02")
	} else if u.Passwd == "" {
		msg2 = "请输入密码 (OvO)"
		fmt.Println("注册：03")
	} else if u.RePasswd != u.Passwd {
		msg2 = "请重新确认您的密码 (?o?)"
		fmt.Println("注册：04")
	} else {
		uname := utils.Trim(utils.Strval(u.Name))
		email := utils.Trim((u.Email))
		pwd := utils.MD5(strings.TrimSpace(u.Passwd))
		// TODO: 数据库查询  <29-12-21, Redari-Es> //

		qs.Name = uname

		fmt.Println("name：", uname)
		qs.Email = email
		fmt.Println("email：", email)
		qs.Passwd = pwd
		fmt.Println("pwd：", u.Passwd, "加密pwd：", pwd)
		// TODO: 数据库插入  <29-12-21, Redari-Es> //
		if err := o.Read(qs, "Name"); err == nil {
			fmt.Println("error", err)
			msg2 = "已有该用户"

		} else if err := o.Read(qs, "Email"); err == nil {
			fmt.Println("error", err)
			msg2 = "已有该邮箱"
		} else {
			o.Insert(qs)
			c.Data["username"] = uname
			c.Data["email"] = email
			c.Data["passwd"] = u.Passwd
			c.Data["repasswd"] = u.RePasswd
			c.Data["md5pwd"] = pwd
			msg = "注册成功，请您登录 !"
			fmt.Println("注册：5 ")
			fmt.Println("注册：注册成功! ")
		}

		//models.Createuser(&u.Name)
		//models.Createemail(&u.Email)
		//	models.Createpwd(&u.Passwd)
		fmt.Println("注册操作结束")
	}
	//	helo := c.GetString("username")
	//	if helo == "" {
	//		c.Ctx.WriteString("hahaha")
	//		return
	//	}

	//	c.Data["errmsg"] = msg
	c.Data["msg"] = msg2
	//c.Redirect("/mall/login", 302)
	c.TplName = "user.html"
	//c.Ctx.Redirect(303, "login")

}
