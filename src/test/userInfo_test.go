package test

import (
	"bookstores2/src/dao"
	"bookstores2/src/model"
	"fmt"
	"testing"
)

var userinfo = &model.UserInfo{
	Username: "li1ng14",
	Password: "lddd123457",
	Email: "l1ng14@foxmail.com",
}

func TestLogin(t *testing.T)  {
	flag ,_,err := dao.Login("lili","123456")
	if err != nil{
		fmt.Println("登录出现异常，err：",err)
	}
	if flag {
		fmt.Println("密码正确，登录成功！")
	}else {
		fmt.Println("密码错误，登录失败。")
	}
}

func TestFindByName(t *testing.T)  {
	info ,err := dao.FindUserByName("lili")
	if err != nil{
		fmt.Println("查询用户信息出现异常，err：",err)
	}else {
		fmt.Println("查询到用户：",info)
	}
}

func TestAddUser(t *testing.T)  {
	err := dao.AddUser(userinfo)
	if err != nil{
		fmt.Println("查询用户信息出现异常，err：",err)
	}else {
		fmt.Println("查询用户，验证是否已经添加：")
		info ,err := dao.FindUserByName(userinfo.Username)
		if err != nil{
			fmt.Println("查询用户信息出现异常，err：",err)
		}else {
			fmt.Println("查询到用户：",info)
		}
	}
}

func TestDeleteUserById(t *testing.T)  {
	_,err := dao.DeleteUserById(10)
	if err != nil{
		fmt.Println("查询用户信息出现异常，err：",err)
	}else {
		fmt.Println("查询用户，验证是否已经删除：")
		info ,err := dao.FindUserById(10)
		if err != nil{
			fmt.Println("查询用户信息出现异常，err：",err)
		}else {
			fmt.Println("查询到用户? ：",info)
		}
	}
}

func TestUpdatePwdById(t *testing.T)  {
	_,err := dao.UpdatePwdById(11,"88888888")
	if err != nil{
		fmt.Println("查询用户信息出现异常，err：",err)
	}else {
		fmt.Println("查询用户，验证是否已经修改密码：")
		info ,err := dao.FindUserById(11)
		if err != nil{
			fmt.Println("查询用户信息出现异常，err：",err)
		}else {
			fmt.Println("查询到用户? ：",info)
		}
	}
}