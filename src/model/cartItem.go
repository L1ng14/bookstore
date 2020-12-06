package model
/*-- 创建购物项表
CREATE TABLE cart_itmes(
id INT PRIMARY KEY AUTO_INCREMENT,
COUNT INT NOT NULL,
amount DOUBLE(11,2) NOT NULL,
book_id INT NOT NULL,
cart_id VARCHAR(100) NOT NULL,
FOREIGN KEY(book_id) REFERENCES books(id),
FOREIGN KEY(cart_id) REFERENCES carts(id)
)*/

type CartItem struct {
	CartItemId int64
	Count      int64
	Amount     float64
	Book       *Book
	CartId     string
}

//获取购物项总价
func (c *CartItem)GetAmount() float64 {
	amount := c.Book.Price * float64(c.Count)
	c.Amount = amount
	return amount
}