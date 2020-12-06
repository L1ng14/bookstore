package model

type Session struct {
	Session_id string
	User_id int
	Username string
	Cart *Cart
	OrderId string
	Orders []*Order
}