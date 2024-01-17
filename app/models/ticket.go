package models

import (
	"registration-booking/global"
	"strconv"
)

type Ticket struct {
	ID
	ScheduleId uint   `json:"schedule_id" gorm:"index:comment:排班ID"`
	Day        string `json:"day" gorm:"index:comment:日期"`
	Num        int    `json:"num" gorm:"comment:可预约人数"`
	Version    int    `json:"version" gorm:"default:1;comment:版本号"`
	Total      int    `json:"total" gorm:"comment:总人数"`
	Timestamps
	SoftDeletes
}

func (ticket Ticket) GetUid() string {
	return strconv.Itoa(int(ticket.ID.ID))
}

func FindTicketsByScheduleId(scheduleId uint, day string) (ticket Ticket, err error) {
	err = global.App.DB.Where("schedule_id = ? AND day = ?", scheduleId, day).Find(&ticket).Error
	return
}

func FindTicketsById(id uint) (ticket Ticket, err error) {
	err = global.App.DB.Where("id = ?", id).Find(&ticket).Error
	return
}
