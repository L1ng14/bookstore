package controller

import (
	"bookstores2/src/dao"
	"bookstores2/src/model"
	"bookstores2/src/utils"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//用户登录处理函数
func Login(w http.ResponseWriter , r *http.Request)  {
	//ToLogin(w ,r)
	username := r.PostFormValue("username")
	success,uid ,err :=dao.Login(username,
							r.PostFormValue("password"))
	if err != nil {
		fmt.Println("登录处理出现错误，err：",err)
	}
	if success{
		//登录成功
		//需要创建session
		uuid := utils.CreateUUID()
		session := &model.Session{
			Session_id: uuid,
			User_id: uid,
			Username: username,
		}
		err := dao.AddSession(session)
		if err != nil{
			fmt.Println("添加session 失败，err:",err)
		}
		//创建一个cookie，并将cookie写入到浏览器中
		cookie := &http.Cookie{
			HttpOnly: true,
			Name: "user",
			Value: uuid,  //未设置过期时间、最长时间，则为会话 cookie，关闭浏览器后失效
		}
		http.SetCookie(w,cookie)

		t :=template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		err = t.Execute(w ,username)
		if err != nil{
			fmt.Fprintln(w, "解析模板出现异常 ，err:",err)
		}
	}else{
		//登录失败处理
		t :=template.Must(template.ParseFiles("views/pages/user/logining.html"))
		err := t.Execute(w ,"登录失败，请检查输入的用户名和密码。")
		if err != nil{
			fmt.Fprintln(w, "解析模板出现异常 ，err:",err)
		}
	}
}

//用户注册处理函数
func Register(w http.ResponseWriter , r *http.Request)  {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	var info = &model.UserInfo{
		Username: username,
		Password: password,
		Email: email,
	}

	row  ,errFind:=dao.FindUserByName(username)

	//验证用户名是否已经存在  (后期改用 AJAX 处理)
	//if row != nil && errFind == nil {
	//	t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
	//	errExe := t.Execute(w, "用户名已存在！请重新输入。")
	//	if errExe != nil {
	//		fmt.Fprintln(w, "解析模板出现异常 ，err:", errExe)
	//	}
	//}

	if row == nil && errFind != nil{
		//进行添加用户操作
		errAdd := dao.AddUser(info)
		if errAdd != nil {
			fmt.Println("注册处理出现错误，err：", errAdd)
		}
		//解析页面模板
		t := template.Must(template.ParseFiles("views/pages/user/logining.html"))
		errExe := t.Execute(w, "")
		if errExe != nil {
			fmt.Fprintln(w, "解析模板出现异常 ，err:", errExe)
		}
	}
}

//通过AJAX 验证用户名是否重复
func FindUserByName(w http.ResponseWriter , r *http.Request){
	row ,err :=dao.FindUserByName(r.PostFormValue("username"))

	if err ==nil && row != nil {
		w.Write([]byte("用户名已存在！请重新输入。"))
	}else {
		w.Write([]byte("<font style='color:blue'>用户名可用。</font>"))
	}
}

//注销当前用户，销毁Cookie
func Logout(w http.ResponseWriter , r *http.Request){
	//获取Cookie 查看是否已经登录
	cookie ,_:= r.Cookie("user")
	if cookie == nil{
		fmt.Println("获取Cookie失败，Cookie可能不存在，还未登录。")
	}else {
		//根据session_id获取Cookie
		session, errFindSession := dao.DeleteSessionById(cookie.Value)
		if session == 0 && errFindSession != nil {
			fmt.Println("数据库中没查找到该session相关记录，err", errFindSession)
		} else {
			//设置为 -1 ，浏览器中的cookie立即销毁
			cookie.MaxAge = -1
			http.SetCookie(w,cookie)
			fmt.Println("刪除相关登录session成功，重定向到主页。")
			http.Redirect(w,r,"/main",302)
		}
	}
}

//登录前预处理，检查是否登录
func ToLogin(w http.ResponseWriter,r *http.Request)  {
	//获取Cookie 查看是否已经登录
	isLogin ,_, _:= dao.IsLogin(r)
	if isLogin {
		//已经登录，重定向到主页
		http.Redirect(w,r,"/main",302)
	}else {
		t :=template.Must(template.ParseFiles("views/pages/user/logining.html"))
		t.Execute(w,"")
	}
}

//首页查找所有图书
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//mapValue := r.URL.Query()
	//pageNo := mapValue.Get("PageNo")
	// ==
	pageNo := r.FormValue("PageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	indexPage, _ := strconv.ParseInt(pageNo, 10, 64)
	pages, err := dao.GetPageBooks(indexPage)
	if err != nil {
		fmt.Println("分页查询全部图书出现异常,err：", err)
	}

	isLogin, session, errFindSession := dao.IsLogin(r)
	if !isLogin || session == nil {
		fmt.Println("数据库中没查找到该session相关记录，err", errFindSession)
		pages.IsLogin = false
	} else {
		pages.IsLogin = true
		pages.Username = session.Username
	}
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, pages)
}
