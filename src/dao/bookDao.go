package dao

import (
	"bookstores2/src/model"
	"bookstores2/src/utils"
	"fmt"
)

//查询所有图书
func FindAllBooks() ( []*model.Book , error)  {
	sqlStr := "select id , title,author ,price ,sales ,stock ,img_path  from books   "
	rows,errQuery  := utils.Db.Query(sqlStr)
	if errQuery != nil{
		fmt.Println("查询多行数据出错，err:",errQuery)
		return nil ,errQuery
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		errNotFound :=rows.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImagePath)
		if errNotFound != nil{
			fmt.Println("赋值结果时出现异常，err :",errNotFound)
			return nil ,errNotFound
		}
		books = append(books, book)
	}
	return  books ,nil
}

//根据图书名查询一条记录
func FindBookByTitle(title string) ( book *model.Book ,err error)  {
	sqlStr := "select id , title,author ,price ,sales ,stock ,img_path  from books where title = ?  "
	row  := utils.Db.QueryRow(sqlStr,title)

	book = &model.Book{}
	errNotFound :=row.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImagePath)
	if errNotFound != nil{
		return nil ,errNotFound
	}
	return  book ,nil
}

//根据图书ID查询一条记录
func FindBookById(id int) ( book *model.Book ,err error)  {
	sqlStr := "select id , title,author ,price ,sales ,stock ,img_path  from books where id = ?  "
	row  := utils.Db.QueryRow(sqlStr,id)

	book = &model.Book{}
	errNotFound :=row.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImagePath)
	if errNotFound != nil{
		return nil ,errNotFound
	}
	return  book ,nil
}

//新增一个图书
func AddBook(book *model.Book) error {
	//事务性，预编译
	sqlStr := "insert into books(title,author ,price ,sales ,stock ,img_path) values (?,?,?,?,?,?);"

	//Prepare创建一个准备好的状态用于之后的查询和命令。
	//返回值可以同时执行多个查询和命令。
	stmt,err := utils.Db.Prepare(sqlStr)
	if err!= nil {
		fmt.Println("预编译 Prepare 出错，err :",err)
		return err
	}
	_ ,errExec := stmt.Exec(&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImagePath)
	if errExec != nil{
		fmt.Println("执行出错,err:",errExec)
		return errExec
	}
	//执行成功！
	return nil
}

//根据图书 ID 删除一个图书
func  DeleteBookById(BookId int) (int64,error){
	sqlStr := "delete from books where id =?"
	stmt ,err := utils.Db.Prepare(sqlStr)
	if err!= nil {
		fmt.Println("预编译 Prepare 出错，err :",err)
		return 0,err
	}
	res ,errExec := stmt.Exec(BookId)
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

//根据图书 更新（先根据id查询出结果，
//并图书所有信息展示在页面，再用查询结果更新数据）
func UpdateBooks(book *model.Book) (int64 ,error){
	sqlStr := "Update books set title = ?,author  = ?,price  = ?,sales  = ?,stock  = ?,img_path  = ? where id = ?"
	stmt ,err := utils.Db.Prepare(sqlStr)
	if err!= nil {
		fmt.Println("预编译 Prepare 出错，err :",err)
		return 0,err
	}
	res ,errExec := stmt.Exec(&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImagePath,&book.ID)
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

//获取总页数，总记录数
func getPage(sqlStr string,pageSize int64,args ...interface{}) (pages int64 ,count int64,err error)   {
	//查询总记录
	res  :=utils.Db.QueryRow(sqlStr,args...)
	err = res.Scan(&count)
	if err != nil{
		fmt.Println("扫描输入 总记录数出现异常，err:",err)
		return 0,0,err
	}
	//计算总页数
	if count % pageSize == 0{
		pages = count / pageSize
	}else {
		pages = count / pageSize +1
	}
	return
}

//为查询到的图书集合赋值
func getBooksForPage(sqlStr string ,args ...interface{})( []*model.Book , error)  {
	rows ,errQuery := utils.Db.Query(sqlStr,args...)
	if errQuery != nil{
		fmt.Println("分页获取部分记录出现异常，err:",errQuery)
		return nil ,errQuery
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		errNotFound :=rows.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImagePath)
		if errNotFound != nil{
			fmt.Println("赋值结果时出现异常，err :",errNotFound)
			return nil ,errNotFound
		}
		books = append(books, book)
	}
	return books,nil
}

//分页查询图书，
//返回的 Page 包含页码，该页包含的图书等信息
func GetPageBooks(IndexPage int64) ( *model.Page , error) {
	//查询总记录
	sqlStr := "select count(*) from books"
	var pageSize int64 = 4

	pages  ,count,err := getPage(sqlStr,pageSize)
	if err !=nil {
		fmt.Println("获取页数页码等出现异常，err:",err)
		return nil, err
	}
	//是否可以传参，不用每次都查询总记录数，减少IO?

	//分页查询
	sqlStr2  := "select * from books limit ? ,?"
	books ,err :=getBooksForPage(sqlStr2,(IndexPage -1) *pageSize,pageSize)
	if err != nil {
		fmt.Println("获取图书集出现异常，err:",err)
		return nil ,err
	}

	page := &model.Page{
		Books: books,
		Pages: pages,
		PageSize: pageSize,
		IndexPage: IndexPage,
		Count: count,
	}
	return page,nil
}

//查询图书价格范围
func QueryPrice(IndexPage int64,low float64,high float64) (  *model.Page , error)  {
	//查询总记录
	sqlStr := "select count(*) from books where price between ? and ?"
	var pageSize int64 = 4

	pages, count, err := getPage(sqlStr, pageSize, low, high)
	if err !=nil {
		fmt.Println("获取页数页码等出现异常，err:",err)
		return nil, err
	}

	sqlStr2 := "select id , title,author ,price ,sales ,stock ,img_path  from books where price between ? and ? limit ? ,?"
	books ,err :=getBooksForPage(sqlStr2,low,high,(IndexPage -1) *pageSize,pageSize)
	if err != nil {
		fmt.Println("获取图书集出现异常，err:",err)
		return nil ,err
	}

	page := &model.Page{
		Books: books,
		Pages: pages,
		PageSize: pageSize,
		IndexPage: IndexPage,
		Count: count,
		MinPrice: low,
		MaxPrice: high,
	}
	return page,nil
}
