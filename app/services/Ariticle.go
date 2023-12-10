package services

import (
	"registration-booking/app/common/request"
	"registration-booking/app/common/response"
	"registration-booking/app/models"
	"strconv"
	"time"
)

type articleService struct{}

var ArticleService = new(articleService)

func (s *articleService) CreateArticle(form request.Article, id string) (error, models.Article) {
	userId, _ := strconv.Atoi(id)
	article := models.Article{
		Title:        form.Title,
		Url:          form.Url,
		Photo:        form.Photo,
		Content:      form.Content,
		DepartmentId: form.DepartmentId,
		Timestamps: models.Timestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			CreatedBy: uint(userId),
			UpdatedBy: uint(userId),
		},
	}

	if err := models.CreateArticle(article); err != nil {
		return err, article
	}
	return nil, article
}

func (s *articleService) GetArticleList(form request.Page) (error, *response.PageData) {
	list, count, err := models.FindAllArticleByPage(form.PageNo, form.PageSize)
	if err != nil {
		return err, nil
	}
	return nil, &response.PageData{
		PageData: list,
		Total:    count,
	}
}

func (s *articleService) GetArticleById(id string) (error, models.Article) {
	idInt, _ := strconv.Atoi(id)
	article, err := models.FindArticleById(uint(idInt))
	if err != nil {
		return err, article
	}
	return nil, article
}

func (s *articleService) UpdateArticle(form request.Article, id string) error {
	idInt, _ := strconv.Atoi(id)
	article, err := models.FindArticleById(form.Id)
	if err != nil {
		return err
	}
	article.Title = form.Title
	article.Url = form.Url
	article.Photo = form.Photo
	article.Content = form.Content
	article.DepartmentId = form.DepartmentId
	article.UpdatedAt = time.Now()
	article.UpdatedBy = uint(idInt)
	if err := models.UpdateArticle(article); err != nil {
		return err
	}
	return nil
}

func (s *articleService) DeleteArticle(id string) error {
	idInt, _ := strconv.Atoi(id)
	if err := models.DeleteArticle(uint(idInt)); err != nil {
		return err
	}
	return nil
}
