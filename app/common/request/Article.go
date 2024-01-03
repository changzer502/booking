package request

type Article struct {
	Id           uint   `form:"id" json:"id"`
	Title        string `form:"title" json:"title"`
	Url          string `form:"url" json:"url"`
	Photo        string `form:"photo" json:"photo"`
	Content      string `form:"content" json:"content"`
	DepartmentId uint   `form:"department_id" json:"department_id"`
}

type GetArticleListReq struct {
	Page
	Dept uint `form:"dept" json:"dept"`
}
