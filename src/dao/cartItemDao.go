package dao

import (
	"bookstores2/src/model"
	"bookstores2/src/utils"
	"fmt"
)

//添加购物项
func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "insert into cart_itmes(COUNT,amount,book_id,cart_id) values(?,?,?,?)"

	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常，err:", err)
		return err
	}
	//购物车项总额通过结构体方法获取
	_, errExec := stmt.Exec(cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartId)
	if errExec != nil {
		fmt.Println("执行出错,err:", errExec)
		return errExec
	}
	//执行成功！
	return nil
}

//根据图书的id和购物车的id获取对应的购物项
func FindCartItemById(bookId int, cartId string) (*model.CartItem, error) {
	sqlStr := "select id, COUNT,amount,book_id,cart_id from cart_itmes where cart_id = ? and book_id = ?"

	cartItem := &model.CartItem{}
	err := utils.Db.QueryRow(sqlStr,cartId,bookId).Scan(&cartItem.CartItemId,&cartItem.Count, &cartItem.Amount,&bookId,&cartItem.CartId)
	if err != nil{
		fmt.Println("查询购物项出现异常，没找到该购物项，err",err)
		return nil ,err
	}
	book ,err := FindBookById(bookId)
	if err != nil {
		fmt.Println("查询图书信息出现异常，err:",err)
	}
	cartItem.Book = book

	return cartItem,nil
}


//根据购物车的id获取购物车中所有的购物项
func FindCartItemsByCartId(cartId string) ([]*model.CartItem,error) {
	sqlStr := "select id ,COUNT,amount,book_id,cart_id from cart_itmes where cart_id = ?"
	var cartItems []*model.CartItem

	rows ,err := utils.Db.Query(sqlStr,cartId)
	if err != nil{
		fmt.Println("查询所有的购物项出现异常，err",err)
		return nil ,err
	}
	for rows.Next(){
		cartItem := &model.CartItem{}
		var bookId int
		//此处直接扫描写入 &cartItem.Book.ID 会panic 因为此时book还未声明，没有空间地址
		err := rows.Scan(&cartItem.CartItemId,&cartItem.Count, &cartItem.Amount,&bookId,&cartItem.CartId)
		if err != nil {
			fmt.Println("扫描写入出现异常，err:",err)
			return nil ,err
		}
		book ,err := FindBookById(bookId)
		if err != nil {
			fmt.Println("查询图书信息出现异常，err:",err)
		}
		cartItem.Book = book
		cartItems = append(cartItems,cartItem)
	}
	return cartItems,nil
}

//更新购物项图书数量,小计金额
func UpdateBookCount(cartItem *model.CartItem) error {
	sqlStr := "update cart_itmes set COUNT = ? , amount = ? where book_id = ? and cart_id = ?"

	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常，err:", err)
		return err
	}
	//购物车项总额通过结构体方法 GetAmount() 获取
	_, errExec := stmt.Exec(cartItem.Count , cartItem.GetAmount(), cartItem.Book.ID,cartItem.CartId)
	if errExec != nil {
		fmt.Println("执行出错,err:", errExec)
		return errExec
	}
	//执行成功！
	return nil
}

//根据购物车id删除购物购物项(配合清空购物车，清空购物车之前，清空购物车内所有购物项)
func DeleteCartItemByCartId(cartId string) error{
	sqlStr := "delete from cart_itmes where cart_id = ?"

	stmt ,err := utils.Db.Prepare(sqlStr)
	if err!= nil {
		fmt.Println("预编译 Prepare 出错，err :",err)
		return err
	}
	_ ,errExec := stmt.Exec(cartId)
	if errExec != nil{
		fmt.Println("执行出错,err:",errExec)
		return errExec
	}
	return  nil
}

//根据购物项id删除购物购物项
func DeleteCartItemByCartItemId(cartItemId int64) error{
	sqlStr := "delete from cart_itmes where id = ?"

	stmt ,err := utils.Db.Prepare(sqlStr)
	if err!= nil {
		fmt.Println("预编译 Prepare 出错，err :",err)
		return err
	}
	_ ,errExec := stmt.Exec(cartItemId)
	if errExec != nil{
		fmt.Println("执行出错,err:",errExec)
		return errExec
	}
	return  nil
}