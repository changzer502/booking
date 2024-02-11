package models

import (
	"registration-booking/global"
	"strconv"
	"strings"
)

type Booking struct {
	ID
	CardId   uint `json:"card_id" gorm:"index:comment:就诊卡ID"`
	TicketId uint `json:"ticket_id" gorm:"index:comment:票ID"`
	UserId   uint `json:"user_id" gorm:"index:comment:用户ID"`
	Status   int  `json:"status" gorm:"index:comment:状态（0未就诊，1已就诊）" default:"0" `
	Rank     int  `json:"rank" gorm:"index:comment:排号"`
	Timestamps
	SoftDeletes
}

func (booking Booking) GetUid() string {
	return strconv.Itoa(int(booking.ID.ID))
}
func FindBookingByTicketIdAndUserId(ticketId uint, userId uint, cardId int) (booking Booking, count int64, err error) {

	err = global.App.DB.Where("ticket_id = ? AND user_id = ? AND card_id = ?", ticketId, userId, cardId).Find(&booking).Count(&count).Error
	return

}
func FindBookingHistoryByUid(userId uint, page, pageSize int) (bookings []Booking, count int64, err error) {
	err = global.App.DB.Where("user_id = ? ", userId).Order("created_at DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&bookings).Count(&count).Error
	return
}
func FindBookingHistoryByDeptId(deptId, doctorId uint, Date string, page, pageSize int, doctorIds, cardIds []string) (bookings []Booking, count int64, err error) {
	query := ""
	if doctorId != 0 {
		query += " AND schedules.doctor_id = " + strconv.Itoa(int(doctorId))
	}
	if len(doctorIds) != 0 {
		query += " AND schedules.doctor_id IN (" + strings.Join(doctorIds, ",") + ")"
	}
	if len(cardIds) != 0 {
		query += " AND bookings.card_id IN (" + strings.Join(cardIds, ",") + ")"
	}
	if Date != "" {
		query += " AND tickets.`day` = '" + Date + "'"
	}

	global.App.DB.Table("bookings").Joins("left join tickets on bookings.ticket_id = tickets.id").Joins("left join schedules on tickets.schedule_id = schedules.id").Where("schedules.department_id = ? "+query, deptId).Order("bookings.rank").Offset((page - 1) * pageSize).Limit(pageSize).Find(&bookings).Count(&count)
	return
}

func FindBookingHistoryById(id uint) (booking Booking, err error) {
	err = global.App.DB.Where("id = ?", id).Find(&booking).Error
	return
}

func FindBookingStatisticsByDept(deptId, doctorId, m string) (count int64, err error) {
	query := ""
	if doctorId != "" {
		query += " AND schedules.doctor_id = " + doctorId
	}
	err = global.App.DB.Table("bookings").Joins("left join tickets on bookings.ticket_id = tickets.id").Joins("left join schedules on tickets.schedule_id = schedules.id").Where("schedules.department_id = ? AND tickets.day LIKE ?"+query, deptId, m+"%").Count(&count).Error
	return
}
