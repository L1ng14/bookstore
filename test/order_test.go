package test

import (
	"bookstores2/src/dao"
	"bookstores2/src/model"
	"bookstores2/src/utils"
	"fmt"
	"testing"
	"time"
)

func TestAddOrder(t *testing.T)  {
	cart, _ := dao.FindCartByUserId(16)
	uuid := utils.CreateUUID()

	order := &model.Order{
		OrderID: uuid,
		CreateTime: time.Now(),
		TotalAmount: cart.GetTotalAmount(),
		TotalCount: cart.GetTotalCount(),
		State: 1,
		UserId: 16,
	}
	dao.AddOrder(order)
}

func TestFindAllOrderByUserId(t *testing.T)  {
	orders ,_:=dao.FindAllOrderByUserId(1)
	for _, v := range orders {
		fmt.Println(v)
	}
}

func TestFindAllOrder(t *testing.T)  {
	orders ,_:=dao.FindAllOrder()
	for _, v := range orders {
		fmt.Println(v)
	}
}

func TestUpdateOrderState(t *testing.T)  {
	dao.UpdateOrderState("23b85b72-1bd3-4805-6ae3-b7c808b6239b",1)
}
