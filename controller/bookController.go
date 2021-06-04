package controller

import (
	"bookstores2/src/dao"
	"bookstores2/src/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//查找所有图书
func FindAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := dao.FindAllBooks()
	if err != nil {
		fmt.Println("查询全部图书出现异常,err：", err)
	}
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}

//分页查找 所有图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//mapValue := r.URL.Query()
	//pageNo := mapValue.Get("PageNo")
	// ==
	pageNo := r.FormValue("PageNo")
	indexPage, _ := strconv.ParseInt(pageNo, 10, 64)
	pages, err := dao.GetPageBooks(indexPage)
	if err != nil {
		fmt.Println("分页查询全部图书出现异常,err：", err)
	}
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, pages)
}

////添加图书
//func AddBook(w http.ResponseWriter ,r *http.Request)  {
//	//获取添加图书的参数
//	title := r.PostFormValue("title")
//	price ,_:= strconv.ParseFloat(r.PostFormValue("price"),64)
//	author := r.PostFormValue("author")
//	sales ,_:= strconv.ParseInt( r.PostFormValue("sales"),10,0)
//	stock ,_ := strconv.ParseInt( r.PostFormValue("stock"),10,0)
//
//	book := &model.Book{
//		Title: title,
//		Author: author,
//		Price: price,
//		Sales: int(sales),
//		Stock: int(stock),
//	}
//	err := dao.AddBook(book)
//	if err != nil {
//		fmt.Println("添加出现异常，err:",err)
//	}
//}

//删除图书
func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	//获取 GET 请求参数
	//mapValue := r.URL.Query()
	//idStr := mapValue.Get("bookId")
	// ==
	idStr := r.FormValue("bookId")

	id, errGetId := strconv.Atoi(idStr)
	if errGetId != nil {
		fmt.Println(w, "获取删除图书的ID出现异常。")
	}

	affect, errDelete := dao.DeleteBookById(id)
	if errDelete != nil {
		fmt.Fprintln(w, "删除图书出现异常,err:", errDelete)
	} else {
		fmt.Println("受影响行数：", affect)
		//执行查找所有图书的处理函数，跳转到图书管理页面
		GetPageBooks(w, r)
	}
}

//跳转到编辑页，更新或添加图书
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	//mapValue := r.URL.Query()
	//idStr := mapValue.Get("bookId")
	// ==
	idStr := r.FormValue("bookId")

	t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))

	if idStr == "" {
		//新增图书
		t.Execute(w, "")
	} else {
		//更新图书处理
		id, errGetId := strconv.Atoi(idStr)
		if errGetId != nil {
			fmt.Println(w, "获取删除图书的ID出现异常。")
		} else {
			book, _ := dao.FindBookById(id)
			t.Execute(w, book)
		}
	}
}

//新增图书，更新图书
func AddOrUpdateBook(w http.ResponseWriter, r *http.Request) {
	//获取表单图书的参数
	id, _ := strconv.Atoi(r.PostFormValue("bookId"))
	title := r.PostFormValue("title")
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	author := r.PostFormValue("author")
	sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
	stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
	book := &model.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Price:  price,
		Sales:  int(sales),
		Stock:  int(stock),
	}
	//id == 0 ，说明没有传值过来，id 为初始值，应执行添加操作
	if id == 0 {
		err := dao.AddBook(book)
		if err != nil {
			fmt.Println("添加出现异常，err:", err)
		}
	} else {
		affect, err := dao.UpdateBooks(book)
		if err != nil {
			fmt.Println("更新 出现异常，err:", err)
		} else {
			fmt.Println("更新操作受影响函数：", affect)
		}
	}
	//执行处理函数，跳转到图书管理页面
	GetPageBooks(w, r)
}

//首页根据价格，范围查找图书
func QueryPrice(w http.ResponseWriter, r *http.Request) {
	/*FormValue返回key为键查询r.Form字段得到结果[]string切片的第一个值。POST和PUT主体中的同名参数 优先 于URL查询字符串。
	如果必要，本函数会隐式调用ParseMultipartForm和ParseForm。*/
	min := r.FormValue("min") //也可以获取 POST 请求表单中的值
	max := r.FormValue("max")
	pageNo := r.FormValue("PageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	indexPage, _ := strconv.ParseInt(pageNo, 10, 64)
	var low float64
	var high float64
	low, _ = strconv.ParseFloat(min, 64) //具体输入限制可以通过JQ实现
	high, _ = strconv.ParseFloat(max, 64)

	pages, err := dao.QueryPrice(indexPage, low, high)
	if err != nil {
		fmt.Println("价格范围查询图书出现异常,err：", err)
	}
	isLogin, session, errFindSession := dao.IsLogin(r)
	if !isLogin || session == nil {
		fmt.Println("数据库中没查找到该session相关记录，err", errFindSession)
		pages.IsLogin = false
	} else {
		pages.IsLogin = true
		pages.Username = session.Username
	}
	pages.MinPrice = low
	pages.MaxPrice = high
	t := template.Must(template.ParseFiles("views/searchOfPrice.html"))
	t.Execute(w, pages)
}
