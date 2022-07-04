package models

import (
	"fmt"
	"mall/models/utils"

	"github.com/astaxie/beego/orm"
)

//登录
//增删改查

//查用户
// TODO: 数据库  <27-12-21, Redari-Es> //
//var o orm.Ormer
//var o = orm.NewOrm()
//var user = new(models.User)
/*
var o orm.Ormer
var user User

func Signin(User string, Pwd string) {

}

/*func Signup(User, Pwd, Repwd, Email string) (bool, bool, error) {
	if Pwd != Repwd {
		//		["errmsg"] = "密码不一样"

		e1 := false
		e2 := false
	}
	user.Name = "test"
	user.Passwd = ""
	err := o.Read(user, user.Name)

	fmt.Println(o.Read(user))
	return e1, e2, err

}
*/
/*
// get user by erify code
func getVerifyUser(user *models.User, code string) bool {
	// use tail hex username query user
	hexStr := code[utils.TimeLimitCodeLength:]
	if b, err := hex.DecodeString(hexStr); err == nil {
		user.UserName = string(b)
		if user.Read("UserName") == nil {
			return true
		}
	}

	return false
}
*/
var o orm.Ormer
var user User

// TODO: 登录  <29-12-21, Redari-Es> //
func CheckUser(uname interface{}) bool {
	b := false
	if o.Read(&uname, "UserName") == nil {
		b = true
		return b
	}
	return b
}
func CheckPwd(pwd string) bool {

	b := false
	pwd = utils.MD5(pwd)
	if o.Read(&pwd, "Passwd") == nil {
		b = true
		return b
	}
	return b
}

// TODO: 注册  <29-12-21, Redari-Es> //
func Hasuser(uname interface{}) bool {
	b := false
	err := o.Read(&uname, "Name")
	if err == nil {
		fmt.Println("已有该用户", err)
	} else {
		fmt.Println("可以使用该用户名")
		b = true
		return b
	}
	return b
}
func Hasemail(email string) bool {
	b := false
	err := o.Read(&email, "Email")
	if err == nil {
		fmt.Println("已有该邮箱", err)
	} else {
		fmt.Println("可以使用该邮箱")
		b = true
		return b
	}
	return b
}

func Createuser(uname interface{}) bool {
	b := false
	id, err := o.Insert(&uname)
	if err == nil {
		fmt.Println("创建用户：", uname)
		fmt.Println("id:", id, "用户为：", uname)
		b = true
	} else {
		fmt.Println("创建失败：")
	}
	return b

}

func Createemail(email string) bool {
	b := false
	id, err := o.Insert(&email)
	if err == nil {
		fmt.Println("用户名插入成功")
		fmt.Println("id:", id, "邮箱为：", email)
		b = true
	} else {
		fmt.Println("创建失败：")
	}
	return b
}

func Createpwd(pwd string) bool {
	b := false
	pwd = utils.MD5(utils.Trim(pwd))
	id, err := o.Insert(&pwd)
	if err == nil {
		fmt.Println("密码插入成功")
		fmt.Println("id:", id, "密码为：", pwd)
		b = true
	} else {
		fmt.Println("创建失败!!!")
	}
	return b
}
func Createuid(uid string) bool {
	b := false
	id, err := o.Insert(&uid)
	if err == nil {
		fmt.Println("id:", id, "插入uid")
		fmt.Println("uid:", uid)
		b = true
	} else {
		fmt.Println("创建失败!!!")
	}
	return b
}
