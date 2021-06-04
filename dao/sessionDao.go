package dao

import (
	"bookstores2/src/model"
	"bookstores2/src/utils"
	"fmt"
	"net/http"
)

//添加一个session，
//session_id使用MD5加密。
func AddSession(s *model.Session) error {
	sqlStr := "insert into sessions(session_id , username , user_id) values (?,?,?)"
	//Prepare创建一个准备好的状态用于之后的查询和命令。
	//返回值可以同时执行多个查询和命令。
	stmt,err := utils.Db.Prepare(sqlStr)
	if err!= nil {
		fmt.Println("预编译 Prepare 出错，err :",err)
		return err
	}
	//MD5 加密密码
	s.Session_id =  utils.Md5(s.Session_id)

	_ ,errExec := stmt.Exec(s.Session_id,s.Username,s.User_id)
	if errExec != nil{
		fmt.Println("执行出错,err:",errExec)
		return errExec
	}
	//执行成功！
	return nil
}

//根据session_id删除一个session，
//MD5加密session_id删除数据库中的相应的session记录。
func DeleteSessionById(session_id string) (int64 ,error){
	sqlStr := "delete from sessions where session_id =?"
	stmt ,err := utils.Db.Prepare(sqlStr)
	if err!= nil {
		fmt.Println("预编译 Prepare 出错，err :",err)
		return 0,err
	}
	//MD5 加密密码
	session_id =  utils.Md5(session_id)

	res ,errExec := stmt.Exec(session_id)
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


//根据session_id查找一个session，用于验证用户是否登录,返回session（包含用户名等信息）
//MD5加密session_id查找数据库中的相应的session记录。
func IsLogin(r *http.Request) ( bool ,*model.Session, error){
	//获取Cookie 查看是否已经登录
	cookie ,_:= r.Cookie("user")
	if cookie == nil{
		fmt.Println("获取Cookie失败，Cookie可能不存在，还未登录。")
		return false,nil,nil
	}else {
		sqlStr := "select session_id,user_id,username from sessions where session_id =?"

		//MD5 加密密码
		session_id := utils.Md5(cookie.Value)

		 res := &model.Session{}
		err  := utils.Db.QueryRow(sqlStr,session_id).Scan(&res.Session_id,&res.User_id,&res.Username)
		if err != nil {
			fmt.Println("数据库中没查找到该session相关记录，err:",err)
			return false,nil,err
		}else {
			return true,res,nil
		}
	}
}
