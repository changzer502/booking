package handler

import (
	"github.com/gin-gonic/gin"
	"registration-booking/app/common/request"
	"registration-booking/app/common/response"
	"registration-booking/app/services"
)

func GetLetterList(c *gin.Context) {
	var form request.Page

	if res, err := services.MessageService.GetLetterList(c.Keys["id"].(string), form.PageNo, form.PageSize); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}
