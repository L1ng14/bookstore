package test

import (
	"bookstores2/src/dao"
	"bookstores2/src/model"
	"bookstores2/src/utils"
	"fmt"
	"testing"
)

func TestAddSession(t *testing.T){
	s := &model.Session{
		Session_id: utils.CreateUUID(),
		Username: "jack",
		User_id: 16,
	}
	err := dao.AddSession(s)
	if err != nil{
		fmt.Println("添加session出现异常,err :",err)
	}
}

//func TestFindSessionById(t *testing.T){
//	sid := "6a3bfd7a-979a-4d1a-50a0-349d6de71a32"
//	res,err := dao.FindSessionById(sid)
//	if err != nil{
//		fmt.Println("添加session出现异常,err :",err)
//	}else {
//		fmt.Println(*res)
//	}
//}

func TestDeleteSession(t *testing.T){
	session_id := "6h78gbx56vc36567fg35htj754kz435c_1604850"
	affect,err := dao.DeleteSessionById(session_id)
	if err != nil{
		fmt.Println("添加session出现异常,err :",err)
	}else{
		fmt.Println("受影响行数为：",affect)
	}
}
