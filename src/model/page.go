package model

type Page struct {
	Pages 		int64		//总页数
	PageSize	int64	//每页显示条数
	Count 		int64		//记录数
	IndexPage 	int64		//当前页
	Books 		[]*Book		//图书结果集
	MinPrice float64
	MaxPrice float64
	IsLogin bool
	Username string
}

func (p *Page)IsHasPrev() bool {
	return p.IndexPage > 1
}

func (p *Page)IsHasNext() bool {
	return p.IndexPage < p.Pages
}

func (p *Page)GetPrevPageNo() int64 {
	return p.IndexPage -1
}

func (p *Page)GetNextPageNo() int64 {
	return p.IndexPage +1
}
