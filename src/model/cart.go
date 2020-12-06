package model

/*
--创建购物车表
CREATE TABLE carts(
id VARCHAR(100) PRIMARY KEY,
total_count INT NOT NULL,
total_amount DOUBLE(11,2) NOT NULL,
user_id INT NOT NULL,
FOREIGN KEY(user_id) REFERENCES users(id)
)*/

type Cart struct {
	CartId      string
	TotalCount  int64
	TotalAmount float64
	UserId      int
	CartItems   []*CartItem
}

//获取图书总数
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _,v := range cart.CartItems {
		totalCount += v.Count
	}
	cart.TotalCount = totalCount
	return totalCount
}

//获取购物车总价
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _,v := range cart.CartItems {
		totalAmount += v.GetAmount()
	}
	cart.TotalAmount = totalAmount
	return totalAmount
}