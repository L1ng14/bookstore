package test

import (
	"bookstores2/src/dao"
	"bookstores2/src/model"
	"fmt"
	"testing"
)

func TestFindCartItemById(t *testing.T)  {
	cartItem ,err := dao.FindCartItemById(2,"bd8f5e14-1ea1-4328-7b8e-845e61ce872f")
	if err != nil {
		fmt.Println("执行FindCartItemById 出现异常，err：",err)
	}
	fmt.Println(*cartItem)
	fmt.Println(*cartItem.Book)
}

func TestFindCartItemsByCartId(t *testing.T){
	cartItems ,err := dao.FindCartItemsByCartId("bd8f5e14-1ea1-4328-7b8e-845e61ce872f")
	if err != nil{
		fmt.Println(err)
	}

	for _,cartItem := range cartItems {
		fmt.Println(*cartItem)
		fmt.Println(*cartItem.Book)
	}
}

func TestUpdateBookCount(t *testing.T){
	cart ,_ := dao.FindCartByUserId(1)
	cartItem ,_ := dao.FindCartItemById(2,cart.CartId)
	if cartItem == nil{
		//未创建该图书相关的购物项
		cartItem := &model.CartItem{
			Count:  1,
			Amount: book.Price,
			Book:   book,
			CartId: cart.CartId,
		}
		dao.AddCartItem(cartItem)
		cart.CartItems = append(cart.CartItems ,cartItem)
	}else {
		//已经创建该购物项 ，购物车内购物项图书 +1 即可
		for _,v := range cart.CartItems{
			if v.Book.ID == cartItem.Book.ID{
				v.Count +=1
				dao.UpdateBookCount(v)
			}
		}
	}
	dao.UpdateCart(cart)
	fmt.Println("   ")
}

func TestDeleteCartItemByCartId(t *testing.T){
	dao.DeleteCartItemByCartId("e0d438a2-24ba-4b6f-623f-697dc2debc5d")
}

func TestDeleteCartItemByCartItemId(t *testing.T){
	dao.DeleteCartItemByCartItemId(20)
}