package test

import (
	"bookstores2/src/dao"
	"bookstores2/src/model"
	"bookstores2/src/utils"
	"fmt"
	"testing"
)

func TestAddCart(t *testing.T)  {
	//设置要买的第一本书
	book := &model.Book{
		ID:    1,
		Price: 27.00,
	}
	//设置要买的第二本书
	book2 := &model.Book{
		ID:    2,
		Price: 23.00,
	}
	//创建一个购物项切片
	var cartItems []*model.CartItem
	cartuuid := utils.CreateUUID()
	//创建两个购物项
	cartItem := &model.CartItem{
		Book:   book,
		Count:  10,
		CartId: cartuuid,
	}
	cartItems = append(cartItems, cartItem)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartId: cartuuid,
	}
	cartItems = append(cartItems, cartItem2)
	//创建购物车
	cart := &model.Cart{
		CartId:    cartuuid,
		CartItems: cartItems,
		UserId:    1,
	}

	 err :=dao.AddCart(cart)
	if err != nil {
		fmt.Println("添加购物车出现异常，err:",err)
	}
}

func TestFindCartByUserId(t *testing.T)  {
	cart ,_:=dao.FindCartByUserId(1)
	fmt.Println(*cart)
	for _,cartItem := range cart.CartItems {
		fmt.Println(*cartItem)
		fmt.Println(*cartItem.Book)
	}
}

func TestUpdateCart(t *testing.T)  {
	cart ,_:=dao.FindCartByUserId(1)
	for _ ,v := range  cart.CartItems{
		v.Count += 1
		dao.UpdateBookCount(v)
	}
	dao.UpdateCart(cart)
}

func TestDeleteCart(t *testing.T){
	dao.DeleteCartById("e0d438a2-24ba-4b6f-623f-697dc2debc5d")
}


