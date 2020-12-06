package controller

import (
	"bookstores2/src/dao"
	"bookstores2/src/model"
	"bookstores2/src/utils"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

//添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//获取Cookie 查看是否已经登录
	isLogin, session, _ := dao.IsLogin(r)

	if isLogin {
		//已经登录
		//获取需要添加的图书 id
		bookId := r.FormValue("bookId")
		bid, _ := strconv.Atoi(bookId)
		book, _ := dao.FindBookById(bid)
		//是否已经创建购物车
		cart, _ := dao.FindCartByUserId(session.User_id)
		if cart == nil {
			//未创建购物车
			cartId := utils.CreateUUID()
			cartItem := &model.CartItem{
				Count:  1,
				Book:   book,
				CartId: cartId,
			}
			var cartItems []*model.CartItem
			cartItems = append(cartItems, cartItem)
			cart = &model.Cart{
				CartId:    cartId,
				UserId:    session.User_id,
				CartItems: cartItems,
			}
			cart.TotalCount = cart.GetTotalCount()
			cart.TotalAmount = cart.GetTotalAmount()
			dao.AddCart(cart)
		} else {
			//已经创建购物车，是否已经创建购物项
			cartItem, _ := dao.FindCartItemById(bid, cart.CartId)
			if cartItem == nil {
				//未创建该图书相关的购物项
				cartItem := &model.CartItem{
					Count:  1,
					Amount: book.Price,
					Book:   book,
					CartId: cart.CartId,
				}
				dao.AddCartItem(cartItem)
				cart.CartItems = append(cart.CartItems, cartItem)
			} else {
				//已经创建该购物项 ，购物车内购物项图书 +1 即可
				for _, v := range cart.CartItems {
					//只有新增图书+1，用cart结构体修改编译后面UpdateCart使用cart
					if v.Book.ID == cartItem.Book.ID {
						v.Count += 1
						dao.UpdateBookCount(v)
					}
				}
			}
			dao.UpdateCart(cart)
		}
		w.Write([]byte("您刚刚将 《" + book.Title + "》 添加到了购物车！"))
	} else {
		//未登录,先进行登录
		w.Write([]byte("请先进行登录再操作！"))
	}
}

//获取当前用户购物车详情
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	isLogin, session, _ := dao.IsLogin(r)
	if isLogin {
		cart, _ := dao.FindCartByUserId(session.User_id)
		session.Cart = cart
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, session)
	} else {
		http.Redirect(w, r, "/toLogin", 302)
	}
}

//清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartId := r.FormValue("cartId")
	//删除购物车之前，删除所有cartId下的购物项
	dao.DeleteCartItemByCartId(cartId)
	dao.DeleteCartById(cartId)

	http.Redirect(w, r, "/getCartInfo", 302)
}

//根据购物项id删除购物购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemId := r.FormValue("cartItemId")
	cId, _ := strconv.ParseInt(cartItemId, 10, 64)
	//根据 cartItemId 删除购物项
	dao.DeleteCartItemByCartItemId(cId)
	//删除之后更新购物车信息
	_, session, _ := dao.IsLogin(r)
	cart, _ := dao.FindCartByUserId(session.User_id)
	cartItems := cart.CartItems
	for i, v := range cartItems {
		//从结构体中移除已经删除的购物项
		if v.CartItemId == cId {
			//截取切片，移除已经删除的购物项
			cartItems = append(cartItems[:i],cartItems[i+1:]...)
			cart.CartItems = cartItems
		}
	}
	//更新购物车
	dao.UpdateCart(cart)

	http.Redirect(w, r, "/getCartInfo", 302)
}

// AJAX 实时更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	bId := r.FormValue("bookId")
	bCount := r.FormValue("bookCount")
	cartId := r.FormValue("cartId")
	bookId, _ := strconv.Atoi(bId)
	bookCount, _ := strconv.ParseInt(bCount, 10, 64)

	_, session, _ := dao.IsLogin(r)
	cart, _ := dao.FindCartByUserId(session.User_id)

	cartItem, _ := dao.FindCartItemById(bookId, cartId)
	//已经创建该购物项 ，购物车内购物项图书数量直接写入即可
	for _, v := range cart.CartItems {
		//只有新增图书+1，用cart结构体修改编译后面UpdateCart使用cart
		if v.Book.ID == cartItem.Book.ID {
			//两个结构体的 Count 都修改，方便后面 Data 传值
			v.Count  =  bookCount
			cartItem.Count = bookCount
			dao.UpdateBookCount(v)
		}
	}
	dao.UpdateCart(cart)

	data := &model.Data{
		TotalAmount: cart.GetTotalAmount(),  //更新购物车获得
		TotalCount: cart.GetTotalCount(),
		Amount: cartItem.GetAmount(),
	}
	//将data转换为json字符串
	json, _ := json.Marshal(data)
	//响应到浏览器
	w.Write(json)
}
