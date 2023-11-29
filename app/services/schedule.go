package services

import (
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
			// 查询预约状态
			if getScheduleListReq.UserId != 0 {
				// TODO 查询预约状态
				scheduleAndTicket.Status = false
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
