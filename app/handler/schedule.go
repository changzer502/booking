package handler

import (
	"github.com/gin-gonic/gin"
	"registration-booking/app/common/request"
	"registration-booking/app/common/response"
	"registration-booking/app/services"
)

func CreateSchedule(c *gin.Context) {
	var form request.CreateScheduleReq
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.ScheduleService.CreateSchedule(form, c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

func CreateTicket(c *gin.Context) {
	var form request.CreateTicketReq
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.ScheduleService.CreateTicket(form, c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

func GetScheduleList(c *gin.Context) {
	var form request.GetScheduleListReq
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.ScheduleService.GetScheduleList(form); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
