package dao

import (
	"bookstores2/src/model"
	"bookstores2/src/utils"
	"fmt"
)

//根据用户名查询一条记录
func FindUserByName(username string) (info *model.UserInfo ,err error)  {
	sqlStr := "select ID ,username ,password ,email from users where username = ?  "
	row  := utils.Db.QueryRow(sqlStr,username)

	info = &model.UserInfo{}
	errNotFound :=row.Scan(&info.ID,&info.Username,&info.Password,&info.Email)
	if errNotFound != nil{
		 return nil ,errNotFound
	}
	return info ,nil
}

//根据用户ID查询一条记录
func FindUserById(id int) (info *model.UserInfo ,err error)  {
	sqlStr := "select ID ,username ,password ,email from users where id = ?  "
	row  := utils.Db.QueryRow(sqlStr,id)

	info = &model.UserInfo{}
	row.Scan(&info.ID,&info.Username,&info.Password,&info.Email)
	return info ,nil
}

//根据用户名查询、验证密码
func Login(username string ,password string) (b bool ,uid int,err error)  {
	sqlStr := "select id,password  from users where username = ? "
	pwd := " "
	b = false

	err  = utils.Db.QueryRow(sqlStr,username).Scan(&uid,&pwd)
	if err != nil{
		return b ,0,err
	}else {
		//MD5加密后比较
		code := utils.Md5(password)
		if code == pwd {  //不区分大小比较
			b = true
		}else{
			return b,0,nil
		}
	}
	return b ,uid,nil
}

//新增一个用户
func AddUser(info *model.UserInfo) error {
	//事务性，预编译
	sqlStr := "insert into users(username,password,email) values (?,?,?);"

	//Prepare创建一个准备好的状态用于之后的查询和命令。
	//返回值可以同时执行多个查询和命令。
	stmt,err := utils.Db.Prepare(sqlStr)
	if err!= nil {
		fmt.Println("预编译 Prepare 出错，err :",err)
		return err
	}
	//MD5 加密密码
	info.Password =  utils.Md5(info.Password)
	_ ,errExec := stmt.Exec(info.Username,info.Password,info.Email)
	if errExec != nil{
		fmt.Println("执行出错,err:",errExec)
		return errExec
	}
	//执行成功！
	return nil
}

//根据用户 ID 删除一个用户
func  DeleteUserById(userId int) (int64,error){
	sqlStr := "delete from users where id =?"
	stmt ,err := utils.Db.Prepare(sqlStr)
	if err!= nil {
		fmt.Println("预编译 Prepare 出错，err :",err)
		return 0,err
	}
	res ,errExec := stmt.Exec(userId)
	if errExec != nil{
		fmt.Println("执行出错,err:",errExec)
		return 0,errExec
	}
	affect ,errRes :=res.RowsAffected()
	if errRes !=nil {
		fmt.Println("取出受影响行数时出现异常，err:",errRes)
		return 0 ,errRes
	}
	return affect ,nil
}

//根据用户 ID 修改用户的密码
func  UpdatePwdById(id int ,password string) ( int64,error){
	str := "update users set password =? where id = ?;"
	stmt ,err := utils.Db.Prepare(str)
	if err!= nil {
		fmt.Println("预编译 Prepare 出错，err :",err)
		return 0,err
	}
	//MD5 加密密码
	password  = utils.Md5(password)
	res ,errExec := stmt.Exec(password,id)
	if errExec != nil{
		fmt.Println("执行出错,err:",errExec)
		return 0,errExec
	}
	affect ,errRes :=res.RowsAffected()
	if errRes !=nil {
		fmt.Println("取出受影响行数时出现异常，err:",errRes)
		return 0 ,errRes
	}
	return affect,nil
}
