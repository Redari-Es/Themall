package controllers

import (
	"fmt"
	"mall/models"
	"mall/models/utils"

	"github.com/astaxie/beego/orm"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) Get() {
	o := orm.NewOrm()
	qs := new(models.User)

	if c.IsLogin {
		fmt.Println("正在访问用户主页", c.Loginuser)
		msg = "正在访问用户主页"
		qs.Name = utils.Strval(c.Loginuser)
		err := o.Read(qs, "Name")

		if err == nil {
			//			fmt.Println(qs.Name, qs.Passwd, qs.Id, qs.Email)
			gender := ""
			phone := ""
			role := ""
			uid := ""

			//		c.Data["uid"] = qs.Uid
			if qs.Uid == "" {
				uid = utils.Uid(qs.Id)
				qs.Uid = uid
				o.Update(qs, "Uid")
				fmt.Println("Uid插入成功：", uid)
			} else {
				uid = qs.Uid

			}
			if qs.Gender == 0 {
				gender = "女"
			} else {
				gender = "男"
			}
			if qs.Phone == "" {
				phone = "尚未填写"
			} else {
				phone = qs.Phone
			}
			if qs.Role == 0 {
				role = "尊贵的用户"
			} else {
				role = "管理员"
			}
			c.Data["id"] = qs.Id
			c.Data["name"] = qs.Name
			c.Data["email"] = qs.Email
			c.Data["pwd"] = qs.Passwd
			c.Data["uid"] = uid
			c.Data["gender"] = gender
			c.Data["phone"] = phone
			c.Data["role"] = role
			c.Data["created"] = qs.Created
			c.Data["updated"] = qs.Updated
			c.Data["loginuser"] = c.Loginuser

		}

	} else {
		msg = "请重新登录"
	}
	c.Data["about"] = "about"
	c.Data["love"] = "/mall"
	c.Data["jingxuan"] = "/mall"
	c.Data["dacu"] = "/mall"
	c.Data["one"] = "/mall"
	c.Data["two"] = "/mall"
	c.Data["three"] = "/mall"
	c.Data["errmsg"] = msg

	c.TplName = "admin/admin.html"
}
