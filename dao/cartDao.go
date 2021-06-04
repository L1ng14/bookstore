package dao

import (
	"bookstores2/src/model"
	"bookstores2/src/utils"
	"fmt"
)

//添加购物车
func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts(id ,total_amount,total_count,user_id) values(?,?,?,?)"

	stmt ,err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常，err:",err)
		return err
	}
	//图书总数，总价 通过结构体的方法获取
	_,errExec :=stmt.Exec(cart.CartId,cart.GetTotalAmount(),cart.GetTotalCount(),cart.UserId)
	if errExec != nil{
		fmt.Println("执行出错,err:",errExec)
		return errExec
	}

	//添加购物项到相应的购物车中(重要)
	for _,v := range cart.CartItems{
		AddCartItem(v)
	}

	//执行成功！
	return nil
}


//根据用戶 ID查询得到购物车
func FindCartByUserId(userId int) (  *model.Cart , error)  {
	sqlStr := "select id, total_amount,total_count,user_id  from carts where user_id = ?  "

	cart := &model.Cart{}
	errNotFound  := utils.Db.QueryRow(sqlStr,userId).Scan(&cart.CartId,&cart.TotalAmount,&cart.TotalCount,&cart.UserId)
	if errNotFound != nil{
		return nil ,errNotFound
	}

	cartItems ,_ := FindCartItemsByCartId(cart.CartId)
	for _,cartItem := range cart.CartItems {
		cartItems = append(cartItems,cartItem)
	}
	cart.CartItems = cartItems
	return  cart ,nil
}

//更新购物车 总金额，图书总数量
func UpdateCart(cart *model.Cart) error{
	//set 后跟多个值用 ,  而不是 and 隔开！！！！！
	sqlStr := "update carts set  total_amount = ? , total_count = ? where id = ?"

	stmt ,err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常，err:",err)
		return err
	}
	//图书总数，总价 通过结构体的方法获取
	_,errExec :=stmt.Exec(cart.GetTotalAmount(),cart.GetTotalCount(),cart.CartId)
	if errExec != nil{
		fmt.Println("执行出错,err:",errExec)
		return errExec
	}
	//执行成功！
	return nil
}

//删除购物车（清空购物车）
//清空之前需要删除全部购物车内的购物项
func DeleteCartById(cartId string) error{
	sqlStr := "delete from carts where id = ?"

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