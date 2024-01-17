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

func GetScheduleListByDept(c *gin.Context) {
	var scheduleReq request.ScheduleReq
	if err := c.ShouldBindJSON(&scheduleReq); err != nil {
		response.Fail(c, request.GetErrorMsg(scheduleReq, err))
		return
	}
	if scheduleList, err := services.ScheduleService.GetScheduleListByDept(scheduleReq); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, scheduleList)
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

func Booking(c *gin.Context) {
	var form request.BookingReq
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.ScheduleService.Booking(form, c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

func GetInfoByTicketId(c *gin.Context) {
	if ticketInfo, err := services.ScheduleService.GetInfoByTicketId(c.Param("ticket_id")); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, ticketInfo)
	}
}

func BookingHistory(c *gin.Context) {
	var form request.Page
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}
	if bookingHistory, err := services.ScheduleService.BookingHistory(form, c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, bookingHistory)
	}
}

func BookingHistoryByDept(c *gin.Context) {
	var form request.BookingHistoryByDeptReq
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}
	if bookingHistory, err := services.ScheduleService.BookingHistoryByDept(form, c.Param("department_id")); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, bookingHistory)
	}
}

func GetBookingHistoryById(c *gin.Context) {
	if ticketInfo, err := services.ScheduleService.GetBookingHistoryById(c.Param("booking_id"), c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, ticketInfo)
	}
}

func GetBookingStatisticsByDept(c *gin.Context) {
	if Statistics, err := services.ScheduleService.GetBookingStatisticsByDept(c.Param("department_id"), c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, Statistics)
	}
}
