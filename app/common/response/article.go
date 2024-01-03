package response

import "registration-booking/app/models"

type ArticleRes struct {
	models.Article
	DeptName string `json:"dept_name"`
}
