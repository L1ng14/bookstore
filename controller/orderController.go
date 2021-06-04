package controller

import (
	"bookstores2/src/dao"
	"bookstores2/src/model"
	"bookstores2/src/utils"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

//结账（新增订单，包含新增订单项）,
//结账成功，清空购物车。
func Checkout(w http.ResponseWriter , r *http.Request)  {
	_, session, _ := dao.IsLogin(r)
	cart, _ := dao.FindCartByUserId(session.User_id)
	uuid := utils.CreateUUID()

	order := &model.Order{
		OrderID: uuid,
		CreateTime: time.Now(),
		TotalAmount: cart.GetTotalAmount(),
		TotalCount: cart.GetTotalCount(),
		State: 0,
		UserId: session.User_id,
	}
	dao.AddOrder(order)


	for _ ,v := range cart.CartItems{
		//持久化订单项
		orderItem := &model.OrderItem{
			Count: v.Count,
			Amount: v.GetAmount(),
			Title: v.Book.Title,
			Author: v.Book.Author,
			Price: v.Book.Price,
			ImaPath: v.Book.ImagePath,
			OrderId: uuid,
		}
		dao.AddOrderItem(orderItem)

		//更新图书库存信息
		v.Book.Stock -= int(v.Count)
		v.Book.Sales += int(v.Count)
		dao.UpdateBooks(v.Book)
	}
	//传递订单号，用于显示
	session.OrderId = uuid

	//结账之后要清空购物车
	dao.DeleteCartItemByCartId(cart.CartId)
	dao.DeleteCartById(cart.CartId)

	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w,session)
}

//获取当前用户的订单
func GetMyOrder(w http.ResponseWriter, r *http.Request)  {
	_, session, _ := dao.IsLogin(r)
	orders, _ := dao.FindAllOrderByUserId(session.User_id)

	session.Orders = orders
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w,session)
}

//获取所有订单，后台管理用
func GetAllOrder(w http.ResponseWriter, r *http.Request)  {
	orders, _ := dao.FindAllOrder()

	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w,orders)
}

//获取订单的详情，可以查看到所有订单项的信息
func GetOrderInfo(w http.ResponseWriter, r *http.Request)  {
	orderId := r.FormValue("orderId")
	orders, _ := dao.FindOrderInfo(orderId)

	t := template.Must(template.ParseFiles("views/pages/order/order_Info.html"))
	t.Execute(w,orders)
}

//后台管理中进行发货，修改订单状态
func SendOrder(w http.ResponseWriter ,r *http.Request)  {
	orderId := r.FormValue("orderId")
	dao.UpdateOrderState(orderId,1)
	GetAllOrder(w,r)
}

//用户退款或退货，修改订单状态
func TakeOrder(w http.ResponseWriter ,r *http.Request)  {
	orderId := r.FormValue("orderId")
	s := r.FormValue("state")
	state ,_:= strconv.ParseInt(s,10,64)

	dao.UpdateOrderState(orderId,state)

	//退货或者取消可以update一下库存。（略）
	GetMyOrder(w,r)
}