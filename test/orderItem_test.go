package test

import (
	"bookstores2/src/dao"
	"bookstores2/src/model"
	"fmt"
	"testing"
)

func TestAddOrderItem(t *testing.T)  {
	cartItem, _ := dao.FindCartItemById(3, "6717d6f2-50b6-4e0c-78b4-9a44f899be29")
	book ,_ := dao.FindBookById(3)

	orderItem := &model.OrderItem{
		Count: cartItem.Count,
		Amount: cartItem.GetAmount(),
		Title: book.Title,
		Author: book.Author,
		Price: book.Price,
		ImaPath: book.ImagePath,
		OrderId: "fdb66f03-ad4d-4816-5d45-e00f515e22ed",
	}
	dao.AddOrderItem(orderItem)
}

func TestFindOrderInfo(t *testing.T)  {
	orderItems, _ := dao.FindOrderInfo("fb975d0b-1183-4ce5-7235-fc9e926db11a")
	for _, v := range orderItems {
		fmt.Println(v)
	}
}
