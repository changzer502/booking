package services

import (
	"errors"
	"registration-booking/app/common/request"
	"registration-booking/app/common/response"
	"registration-booking/app/models"
	"registration-booking/global"
	"strconv"
	"time"
)

type scheduleService struct {
}

var ScheduleService = new(scheduleService)

func (scheduleService *scheduleService) CreateSchedule(params request.CreateScheduleReq, id string) (err error, user models.Schedule) {
	uid, _ := strconv.Atoi(id)
	user = models.Schedule{
		DoctorId:     params.DoctorId,
		Time:         params.Time,
		Week:         params.Week,
		DepartmentId: params.DepartmentId,
		Price:        params.Price,
		Timestamps: models.Timestamps{
			CreatedBy: uint(uid),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UpdatedBy: uint(uid),
		},
	}
	err = global.App.DB.Create(&user).Error
	return
}

func (scheduleService *scheduleService) GetScheduleList(getScheduleListReq request.GetScheduleListReq) (err error, scheduleList []response.ScheduleList) {
	// 查询科室下的排班医生
	schedules, err := models.FindSchedulesByDepartmentID(getScheduleListReq.DepartmentId, getScheduleListReq.Week)
	if err != nil {
		return err, nil
	}
	doctorMap := make(map[uint][]models.Schedule, 0)
	for i := 0; i < len(schedules); i++ {
		doctorMap[schedules[i].DoctorId] = append(doctorMap[schedules[i].DoctorId], schedules[i])
	}
	for doctorId, schedule := range doctorMap {
		doctor, err := models.FindDoctorById(doctorId)
		if err != nil {
			return err, nil
		}
		scheduleAndTicketLst := make([]response.ScheduleAndTicket, 0)
		for i := 0; i < len(schedule); i++ {
			scheduleAndTicket := response.ScheduleAndTicket{}
			scheduleAndTicket.Schedule = schedule[i]
			// 查询票数
			ticket, err := models.FindTicketsByScheduleId(schedule[i].ID.ID, getScheduleListReq.Day)
			if err != nil {
				return err, nil
			}
			scheduleAndTicket.Ticket = ticket

			if getScheduleListReq.UserId != 0 {
				// 查询预约状态
				_, count, err := models.FindBookingByTicketIdAndUserId(ticket.ID.ID, getScheduleListReq.UserId, getScheduleListReq.CardId)
				if err != nil {
					return err, nil
				}
				if count > 0 {
					scheduleAndTicket.Status = true
				}
			}
			scheduleAndTicketLst = append(scheduleAndTicketLst, scheduleAndTicket)
		}
		add := false
		if getScheduleListReq.Only {
			for i := 0; i < len(scheduleAndTicketLst); i++ {
				if scheduleAndTicketLst[i].Ticket.Num > 0 {
					add = true
				}
			}
		}
		if !getScheduleListReq.Only || add {
			scheduleList = append(scheduleList, response.ScheduleList{
				Doctor:                doctor,
				ScheduleAndTicketList: scheduleAndTicketLst,
			})
		}
	}

	return
}

func (scheduleService *scheduleService) CreateTicket(params request.CreateTicketReq, id string) (err error, ticket models.Ticket) {
	uid, _ := strconv.Atoi(id)
	ticket = models.Ticket{
		ScheduleId: params.ScheduleId,
		Day:        params.Day,
		Num:        params.Num,
		Version:    1,
		Timestamps: models.Timestamps{
			CreatedBy: uint(uid),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UpdatedBy: uint(uid),
		},
	}
	err = global.App.DB.Create(&ticket).Error
	return
}
func (scheduleService *scheduleService) Booking(getScheduleListReq request.BookingReq, id string) (err error, scheduleList []response.ScheduleList) {
	uid, _ := strconv.Atoi(id)
	// 查询票数
	ticket, err := models.FindTicketsById(getScheduleListReq.TicketId)
	if err != nil {
		return err, nil
	}
	if ticket.Num <= 0 {
		return errors.New("暂无"), nil
	}
	// 查询是否已经预约
	_, count, err := models.FindBookingByTicketIdAndUserId(getScheduleListReq.TicketId, uint(uid), int(getScheduleListReq.CardId))
	if err != nil {
		return err, nil
	}
	if count > 0 {
		return errors.New("已预约"), nil
	}
	// 事务开始后，需要使用 tx 处理数据
	tx := global.App.DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	// 减少票数
	ticket.Num = ticket.Num - 1
	version := ticket.Version
	ticket.Version = ticket.Version + 1
	affected := tx.Model(&ticket).Where("id = ? AND version = ?", ticket.ID.ID, version).Updates(&ticket).RowsAffected
	if affected == 0 {
		return errors.New("服务忙请重试~"), nil
	}
	// 创建预约
	booking := models.Booking{
		CardId:   getScheduleListReq.CardId,
		TicketId: ticket.ID.ID,
		UserId:   uint(uid),
	}
	tx.Create(&booking)
	return
}

func (scheduleService *scheduleService) GetInfoByTicketId(ticketId string) (ticketInfo response.TicketInfo, err error) {
	id, _ := strconv.Atoi(ticketId)
	ticket, err := models.FindTicketsById(uint(id))
	if err != nil {
		return
	}
	ticketInfo.Ticket = ticket
	schedule, err := models.FindScheduleByID(ticket.ScheduleId)
	if err != nil {
		return
	}
	ticketInfo.Schedule = schedule
	doctor, err := models.FindDoctorById(schedule.DoctorId)
	if err != nil {
		return
	}
	ticketInfo.Doctor = doctor
	return
}
func (scheduleService *scheduleService) BookingHistory(page request.Page, id string) (res *response.PageData, err error) {
	res = &response.PageData{}
	uid, _ := strconv.Atoi(id)
	bookings, count, err := models.FindBookingHistoryByUid(uint(uid), page.PageNo, page.PageSize)
	if err != nil {
		return nil, err
	}
	bookingInfs := make([]response.BookingInfo, 0)
	for i := 0; i < len(bookings); i++ {
		ticket, err := models.FindTicketsById(bookings[i].TicketId)
		if err != nil {
			return nil, err
		}
		schedule, err := models.FindScheduleByID(ticket.ScheduleId)
		if err != nil {
			return nil, err
		}
		doctor, err := models.FindDoctorById(schedule.DoctorId)
		if err != nil {
			return nil, err
		}
		department, err := models.FindDepartmentById(schedule.DepartmentId)
		if err != nil {
			return nil, err
		}
		card, err := models.FindCardById(bookings[i].CardId)
		if err != nil {
			return nil, err
		}
		bookingInfs = append(bookingInfs, response.BookingInfo{
			Doctor:     doctor,
			Schedule:   schedule,
			Ticket:     ticket,
			Card:       card,
			Department: department,
			Booking:    bookings[i],
		})
	}
	res.PageData = bookingInfs
	res.Total = count
	return
}

func (scheduleService *scheduleService) GetBookingHistoryById(bookingId string, id string) (bookingInfo *response.BookingInfo, err error) {
	bookingIdInt, _ := strconv.Atoi(bookingId)
	booking, err := models.FindBookingHistoryById(uint(bookingIdInt))
	if err != nil {
		return nil, err
	}
	ticket, err := models.FindTicketsById(booking.TicketId)
	if err != nil {
		return nil, err
	}
	schedule, err := models.FindScheduleByID(ticket.ScheduleId)
	if err != nil {
		return nil, err
	}
	doctor, err := models.FindDoctorById(schedule.DoctorId)
	if err != nil {
		return nil, err
	}
	department, err := models.FindDepartmentById(schedule.DepartmentId)
	if err != nil {
		return nil, err
	}
	card, err := models.FindCardById(booking.CardId)
	if err != nil {
		return nil, err
	}
	bookingInfo = &response.BookingInfo{
		Doctor:     doctor,
		Schedule:   schedule,
		Ticket:     ticket,
		Card:       card,
		Department: department,
	}
	return
}
