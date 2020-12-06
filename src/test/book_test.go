package test

import (
	"bookstores2/src/dao"
	"bookstores2/src/model"
	"fmt"
	"testing"
)

var book = &model.Book{
	Title: "Mysql必知必会",
	Author: "Ben Forta",
	Price: 33,
	Sales: 99,
	Stock: 100,
	ImagePath: "static/img/default.jpg",
}

func TestFindAllBooks(t *testing.T)  {
	books ,err := dao.FindAllBooks()
	if err != nil{
		fmt.Println("查询全部图书出现异常,err：",err)
	}
	for v := range books {
		fmt.Println(v)
	}
}

func TestAddBook(t *testing.T)  {
	err := dao.AddBook(book)
	if err !=nil{
		fmt.Println("添加图书出现异常,err :",err)
	}
}

func TestFindBookById(t *testing.T)  {
	bookRes ,err := dao.FindBookById(2)
	if err != nil{
		fmt.Println("根据ID查询出现异常，err：",err)
	}else {
		fmt.Println("根据ID查询，结果为：",*bookRes)
	}
}

func TestFindBookByTitle(t *testing.T)  {
	bookRes ,err := dao.FindBookByTitle("边城")
	if err != nil{
		fmt.Println("根据图书名查询出现异常，err：",err)
	}else {
		fmt.Println("根据图书名查询，结果为：",*bookRes)
	}
}

func TestDeleteBookById(t *testing.T)  {
	res ,err := dao.DeleteBookById(31)
	if err != nil{
		fmt.Println("根据图书ID删除出现异常，err：",err)
	}else {
		fmt.Println("根据图书ID删除，受影响行数为：",res)
	}
}

func TestUpdateBooks(t *testing.T)  {
	bookRes ,err := dao.FindBookById(37)
	if err != nil {
		fmt.Println("根据ID查询图书出现异常，err:",err)
	}
	bookRes.Title = "必知必会-2"
	res ,err := dao.UpdateBooks(bookRes)
	if err != nil{
		fmt.Println("根据图书更新出现异常，err：",err)
	}else {
		fmt.Println("根据图书更新，受影响行数为：",res)
	}
}

func TestGetPageBooks(t *testing.T)  {
	page ,err := dao.GetPageBooks(6)
	if err != nil{
		fmt.Println("分页 查询图书出现异常,err：",err)
	}
	for _ ,v:= range page.Books {
		fmt.Println(*v)
	}
}

func TestQueryPrice(t *testing.T)  {
	pages ,err := dao.QueryPrice(2,30,40)
	if err != nil{
		fmt.Println("分页 查询图书出现异常,err：",err)
	}
	for _ ,v:= range pages.Books {
		fmt.Println(*v)
	}
}
