package model

/*CREATE TABLE order_items(
id INT PRIMARY KEY AUTO_INCREMENT,
COUNT INT NOT NULL,
amount DOUBLE(11,2) NOT NULL,
title VARCHAR(100) NOT NULL,
author VARCHAR(100) NOT NULL,
price DOUBLE(11,2) NOT NULL,
img_path VARCHAR(100) NOT NULL,
order_id VARCHAR(100) NOT NULL,
FOREIGN KEY(order_id) REFERENCES orders(id)
)*/

type OrderItem struct {
	ID int
	Count int64
	Amount float64
	Title string
	Author string
	Price float64
	ImaPath string
	OrderId string
}


////在购物车中已经计算完成，传值即可
//func (o *OrderItem)GetAmount() float64 {
//	var amount float64
//	amount = float64(o.Count) * o.Price
//	o.Amount = amount
//	return amount
//}
