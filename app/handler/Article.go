package handler

import (
	"github.com/gin-gonic/gin"
	"registration-booking/app/common/request"
	"registration-booking/app/common/response"
	"registration-booking/app/services"
)

func CreateArticle(c *gin.Context) {
	var form request.Article
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.ArticleService.CreateArticle(form, c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

func GetArticleList(c *gin.Context) {
	var form request.Page
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, list := services.ArticleService.GetArticleList(form); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, list)
	}
}

func GetArticleById(c *gin.Context) {
	if err, card := services.ArticleService.GetArticleById(c.Param("id")); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, card)
	}
}

func UpdateArticle(c *gin.Context) {
	var form request.Article
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err := services.ArticleService.UpdateArticle(form, c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, nil)
	}
}

func DeleteArticle(c *gin.Context) {
	if err := services.ArticleService.DeleteArticle(c.Param("id")); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, nil)
	}
}
