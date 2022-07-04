/* TODO: 购物车  <23-12-21, Redari-Es> */
//定义存储数据的容器
var datas = {
	p001: { pname: "精选粉底液", price: 1999, src: "/static/mall/pic/5.jpg" },
	p002: { pname: "精选口红", price: 2999, src: "/static/mall/pic/5.jpg" },
	p003: { pname: "精选眼影", price: 3999, src: "/static/mall/pic/5.jpg" },
	p004: { pname: "限时优惠", price: 4999, src: "/static/mall/pic/5.jpg" },
};
//点击购买
$(".buy").click(function () {
	//获取data-pid属性的值
	var pid = $(this).attr("data-pid");
	//点击某个数据的详细信息
	var info = datas[pid];
	//判断是否存在   默认不存在
	var isExist = false;
	//遍历
	var index = 0; //匹配数据䣌索引

	$("tbody>tr").each(function (idx, ele) {
		//判断
		if (pid == $(ele).attr("data-pid")) {
			//有重复的
			isExist = true;
			index = idx;
		}
	});
	//alert(isExist)
	if (isExist) {
		//触发某个重复元素的+号的点击事件   gt  大于索引  lt    trigger触发器
		$("tbody>tr").eq(index).find(".add").trigger("click");
	} else {
		//判断没有重复数据䣌
		var html1 = $("#temp").html();
		//替换值
		html1 = html1.replace("{{pid}}", pid);
		html1 = html1.replace("{{src}}", info.src);
		html1 = html1.replace("{{price}}", info.price);
		html1 = html1.replace("{{price}}", info.price);
		html1 = html1.replace("{{pname}}", info.pname);
		//拼接
		$("tbody").append($(html1));
	}
	TotalPrice();
});

//删除
/* $('table').on('click','.del',delete);
				function delete(){
					
				} */
$("table").on("click", ".del", function () {
	//弹框  提醒是否删除
	var flag = confirm("东西这么六币  不要了吗？");
	if (flag) {
		$(this).parent().parent().remove();
	}
	TotalPrice();
});

//3  计算数量和总价
function TotalPrice() {
	var sum = 0; //总价的变量
	var count = 0; //数量变量
	//遍历
	$("tbody>tr").each(function (idx, ele) {
		//判断框是否被选中
		var flag = $(ele).find(".pid")[0].checked;
		//alert(flag);
		if (flag) {
			sum += parseInt($(ele).find(".count").html().substring(1));
			count += parseInt($(ele).find(".number").val());
		}
	});
	//赋值
	$(".total-money").html("&yen;" + sum);
	$(".total").html(count);
}

//减法
$("table").on("click", ".sub", function () {
	//得到input元素里的值
	var num = $(this).next().val();
	if (num == 1) {
		//判断删除
		var flag = confirm("东西这么六币  不要了吗？");
		if (flag) {
			$(this).parent().parent().remove();
		}
	} else {
		num--;
		//更新赋值
		$(this).next().val(num);
		//单价  price
		var price = $(this).parent().siblings(".price").html().substring(1);
		var result = calPrice(num, price);
		//更新赋值
		$(this)
			.parent()
			.next()
			.html("&yen;" + result);
	}
	TotalPrice();
});
//小计
function calPrice(num, price) {
	return num * price;
}

//加法
$("table").on("click", ".add", function () {
	var num = $(this).prev().val();
	num++;
	//更新赋值
	$(this).prev().val(num);
	var price = $(this).parent().prev().html().substring(1);
	//小计
	var result = calPrice(num, price);
	$(this)
		.parent()
		.next()
		.html("&yen;" + result);
	TotalPrice();
});
// 滚动条
/*使用JavaScript来实现*/

var menuHeight = document.getElementById("menu");

var screenHeight = window.innerHeight; //浏览器窗口的内部高度
/* var  screenHeight =  document.documentElement.clientHeight; */
menuHeight.style.height = screenHeight - 80 + "px";

/*使用jQuery来实现*/

/****方法一****/
$(".menu").height($(window).height() - 80);

/****方法二****/
$(".menu").css("height", $(window).height() - 80);
