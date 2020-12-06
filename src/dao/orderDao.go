package dao

import (
	"bookstores2/src/model"
	"bookstores2/src/utils"
	"fmt"
)

//結賬，新增订单，"不"包含新增订单项
func AddOrder(order *model.Order) error {
	sqlStr := "insert into orders(id, create_time, total_count, total_amount, state, user_id) values(?,?,?,?,?,?)"

	stmt ,err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常，err:",err)
		return err
	}

	 _ ,errExec := stmt.Exec(order.OrderID,order.CreateTime,order.TotalCount,order.TotalAmount,
	 						 order.State,order.UserId)
	if errExec != nil {
		fmt.Println("添加订单出现异常，err:",errExec)
		return errExec
	}
	return nil
}

//根据用户id查询该用户的所有订单
func FindAllOrderByUserId(userId int)( []*model.Order ,error) {
	sqlStr := "select id, create_time, total_count, total_amount, state,user_id from orders where user_id = ?"

	rows,err := utils.Db.Query(sqlStr,userId)
	if err != nil {
		fmt.Println("查询用户订单出现异常，err：",err)
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		err := rows.Scan(&order.OrderID,&order.CreateTime,&order.TotalCount,&order.TotalAmount,&order.State,&order.UserId)
		if err != nil {
			fmt.Println("扫描写入出现异常，err:",err)
			return orders,err
		}
		orders = append(orders , order)
	}

	return orders,nil
}

//查询所有订单
func FindAllOrder()( []*model.Order ,error) {
	sqlStr := "select id, create_time, total_count, total_amount, state,user_id from orders "

	rows,err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println("查询用户订单出现异常，err：",err)
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		err := rows.Scan(&order.OrderID,&order.CreateTime,&order.TotalCount,&order.TotalAmount,&order.State,&order.UserId)
		if err != nil {
			fmt.Println("扫描写入出现异常，err:",err)
			return orders,err
		}
		orders = append(orders , order)
	}
	return orders,nil
}

//更新订单的状态信息
func UpdateOrderState(orderId string,state int64) error{
	sqlStr := "update orders set state = ? where id = ?"

	stmt ,err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常，err:",err)
		return err
	}

	_ ,errExec := stmt.Exec(state,orderId)
	if errExec != nil {
		fmt.Println("添加订单出现异常，err:",errExec)
		return errExec
	}
	return nil
}