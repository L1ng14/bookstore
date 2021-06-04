package model

import "time"

/*CREATE TABLE orders(
id VARCHAR(100) PRIMARY KEY,
create_time DATETIME NOT NULL,
total_count INT NOT NULL,
total_amount DOUBLE(11,2) NOT NULL,
state INT NOT NULL,
user_id INT,
FOREIGN KEY(user_id) REFERENCES users(id)
)*/

type Order struct {
	OrderID     string
	CreateTime  time.Time
	TotalCount  int64
	TotalAmount float64
	OrderItems  []*OrderItem
	State       int64 //  0 未发货,  1  已发货 , 2 交易完成 ,  -1 取消（退款退货）
	UserId      int
}

func (o *Order) SendComplate() bool {
	return o.State == 1
}

func (o *Order) NoSend() bool {
	return o.State == 0
}

func (o *Order) Complate() bool {
	return o.State == 2
}
func (o *Order) Cancel() bool {
	return o.State == -1
}
