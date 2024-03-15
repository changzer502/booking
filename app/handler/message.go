package handler

import (
	"github.com/gin-gonic/gin"
	"registration-booking/app/common/request"
	"registration-booking/app/common/response"
	"registration-booking/app/services"
)

func GetLetterList(c *gin.Context) {
	var form request.Page
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}
	if res, err := services.MessageService.GetLetterList(c.Keys["id"].(string), form.PageNo, form.PageSize); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}
func UnreadCount(c *gin.Context) {
	if res, err := services.MessageService.UnreadCount(c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func GetConversationDetail(c *gin.Context) {
	var form request.Page
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}
	if res, err := services.MessageService.GetConversationDetail(c.Keys["id"].(string), c.Param("conversationId"), form.PageNo, form.PageSize); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}

func SendLetter(c *gin.Context) {
	var sendMessageReq request.SendMessageReq
	if err := c.ShouldBindJSON(&sendMessageReq); err != nil {
		response.Fail(c, request.GetErrorMsg(sendMessageReq, err))
		return
	}
	if err := services.MessageService.SendLetter(c.Keys["id"].(string), sendMessageReq.ToUserId, sendMessageReq.Content); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, nil)
	}
}

func SendNotice(c *gin.Context) {
	var sendMessageReq request.SendNoticeReq
	if err := c.ShouldBindJSON(&sendMessageReq); err != nil {
		response.Fail(c, request.GetErrorMsg(sendMessageReq, err))
		return
	}
	if err := services.MessageService.SendNotice("1", sendMessageReq.ToUserId, sendMessageReq.Content); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, nil)
	}
}

func GetNoticeList(c *gin.Context) {
	var form request.Page
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}
	if res, _, err := services.MessageService.GetNoticeList(c.Keys["id"].(string), form.PageNo, form.PageSize); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}
