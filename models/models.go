package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//关系映射
/*
----------------------
rel/reverse
RelOneToOne:
type User struct
Profile *Profile `orm:"null;rel(one);on_delete(set_null)"`
----------------------
RelReverseOne:
type Profile struct
User *User `orm:"reverse(one)"`
----------------------
RelForeignKey:
type Post struct
User *User `orm:"rel(fk)"`
----------------------
RelReverseMany:
Post []*Post `orm:"reverse(many)"`fk的反向关系
----------------------
RelManyToMany:
Tags []*Tag `orm:"rel(m2m)"` manytomany relation
----------------------
RelReverseMany:
Post []*Post `orm:"reverse(many)"`

----------------------
rel_table/rel_through
针对orm:"rel(m2m)"


m2m
o2o
*/

type User struct {
	Id      int       `orm:"pk;auto"`
	Uid     string    `orm:"column(uid);size(20)"`
	Name    string    `orm:"column(name);size(20);unique;index"`
	Passwd  string    `orm:"column(passwd);size(50)"`
	Gender  int       `orm:"column(gender);size(3);default(0)"`
	Email   string    `orm:"column(email);size(30)"`
	Phone   string    `orm:"column(phone);size(20);default(0)"`
	Role    int       `orm:"column(role);size(2);"`
	Created time.Time `orm:"column(created);auto_now_add;type(datetime)"`
	Updated time.Time `orm:"column(updated);auto_now;type(datetime)"`
	//表关系
	Shipping  []*Shipping  `orm:"rel(m2m)"`
	Cart      []*Cart      `orm:"rel(m2m)"`
	Order     []*Order     `orm:"rel(m2m)"`
	OrderItem []*OrderItem `orm:"rel(m2m)"`
	PayInfo   []*PayInfo   `orm:"rel(m2m)"`
}
type Shipping struct {
	Id              int       `orm:"pk;auto"`
	Uid             string    `orm:"column(uid);size(20)"`
	ReceiveName     string    `orm:"size(20)"`
	ReceivePhone    string    `orm:"size(20)"`
	ReceiveMobile   string    `orm:"size(20)"`
	ReceiveProvice  string    `orm:"size(20)"`
	ReceiveCity     string    `orm:"size(20)"`
	ReceiveDistrict string    `orm:"size(20)"`
	ReceiveAddress  string    `orm:"size(20)"`
	ReceiveZip      string    `orm:"size(20)"`
	Created         time.Time `orm:"column(created);auto_now_add;type(datetime)"`
	Updated         time.Time `orm:"column(updated);auto_now;type(datetime)"`
	//表关系
	User []*User `orm:"reverse(many)"`
}
type Category struct {
	Id       int       `orm:"pk;auto"`
	ParentId int       `orm:"size(20)"`
	Name     string    `orm:"size(20)"`
	status   int       `orm:"size(10)"`
	SortOder int       `orm:"size(20)"`
	Created  time.Time `orm:"column(created);auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"column(updated);auto_now;type(datetime)"`
}
type Cart struct {
	Id       int       `orm:"pk;auto"`
	Uid      string    `orm:"column(uid);size(20)"`
	ParentId int       `orm:"column(parent_id);size(20)"`
	Quantity int       `orm:"column(quantity)"`
	Price    int       `orm:"column(price);size(20)"`
	Checked  int       `orm:"column(checked);size(2);default(1)"`
	Created  time.Time `orm:"column(created);auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"column(updated);auto_now;type(datetime)"`
	//表关系
	User    []*User    `orm:"reverse(many)"`
	Product []*Product `orm:"rel(m2m)"`
}
type Product struct {
	Id         int        `orm:"pk;auto"`
	CategoryId int        `orm:"column(category_id);size(20)"`
	Name       string     `orm:"column(name);size(40)"`
	Subtitle   string     `orm:"column(subtitle);size(50)"`
	MainImage  string     `orm:"column(main_image);size(50)"`
	SubImages  string     `orm:"column(sub_images);size(50)"`
	Detail     string     `orm:"column(detail);size(50)"`
	Price      int        `orm:"column(price);size(20)"`
	Stock      int        `orm:"column(stock);size(4);default(1)"`
	Status     int        `orm:"column(status);size(4);default(1)"`
	Created    time.Time  `orm:"column(created);auto_now_add;type(datetime)"`
	Updated    time.Time  `orm:"column(updated);auto_now;type(datetime)"`
	OrderItem  *OrderItem `orm:"null;reverse(one)"`
	Cart       []*Cart    `orm:"reverse(many)"`
}
type Order struct {
	Id          int    `orm:"pk;auto"`
	Uid         string `orm:"column(uid);size(20)"`
	OrderNo     int    `orm:"column(order_no);size(20)"`
	ShippingId  int    `orm:"column(shipping_id);size(20)"`
	Payment     int    `orm:"column(payment);size(4);default(1)"`
	PaymentType int    `orm:"column(payment_type);size(4)"`
	Postage     int    `orm:"column(postage);size(4)"`
	//待支付钱
	Status int `orm:"column(status);size(5)"`
	//0取消10未付款20已付款40已发货50交易成功60交易关闭
	PaymentTime time.Time `orm:"column(pay_time);type(datetime)"`
	SendTime    time.Time `orm:"column(send_time);type(datetime)"`
	EndTime     time.Time `orm:"column(end_time);type(datetime)"`
	CloseTime   time.Time `orm:"column(close_time);type(datetime)"`
	//
	Created time.Time `orm:"column(created);auto_now_add;type(datetime)"`
	Updated time.Time `orm:"column(updated);auto_now;type(datetime)"`
	//表关系
	User []*User `orm:"reverse(many)"`
}
type OrderItem struct {
	Id               int       `orm:"pk;auto"`
	Uid              string    `orm:"column(uid);size(20)"`
	OrderNo          int       `orm:"column(order_no);size(20)"`
	ProductName      string    `orm:"column(product_name);size(20)"`
	ProductImage     string    `orm:"column(product_image);size(20)"`
	CurrentUnitPrice int       `orm:"column(current_unit_price);size(20)"`
	Quantity         int       `orm:"column(quantity);size(20)"`
	TotalPrice       int       `orm:"column(total_price);size(20)"`
	Created          time.Time `orm:"column(created);auto_now_add;type(datetime)"`
	Updated          time.Time `orm:"column(updated);auto_now;type(datetime)"`
	//表关系
	User []*User `orm:"reverse(many)"`
	//
	ProductId *Product `orm:"null;rel(one);on_delete(set_null)"`
}
type PayInfo struct {
	Id        int       `orm:"pk;auto"`
	Uid       string    `orm:"column(uid);size(20)"`
	OrderNo   int       `orm:"size(20)"`
	PayStatus int       `orm:"size(4);default(1)"`
	Created   time.Time `orm:"column(created);auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"column(updated);auto_now;type(datetime)"`
	//表关系
	User []*User `orm:"reverse(many)"`
}

func RegisterDB() {
	//注册medel
	orm.RegisterModel(new(User), new(Shipping), new(Category),
		new(Product), new(Cart), new(Order), new(OrderItem), new(PayInfo))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	//err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/mall?charset=utf8", 30)
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/mall?charset=utf8", 30)
	if err != nil {
		beego.Info("连接数据库出错")
		return
	}
	beego.Info("连接数据库成功")
	//设置默认时间
	orm.DefaultTimeLoc = time.UTC
	//自动建表 需要将false改为true
	orm.RunSyncdb("default", false, true)
	//orm.RunSyncdb("default", true, true)
	//初始化ormer
	//	o := orm.NewOrm
	//初始化数据
	//datainit()
}

/*

只能用户注册表
支付账单表
购物车表


*/
