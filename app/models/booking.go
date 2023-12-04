package models

import (
	"registration-booking/global"
	"strconv"
)

type Booking struct {
	ID
	CardId   uint `json:"card_id" gorm:"index:comment:就诊卡ID"`
	TicketId uint `json:"ticket_id" gorm:"index:comment:票ID"`
	UserId   uint `json:"user_id" gorm:"index:comment:用户ID"`
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
	err = global.App.DB.Where("user_id = ? ", userId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&bookings).Count(&count).Error
	return
}

func FindBookingHistoryById(id uint) (booking Booking, err error) {
	err = global.App.DB.Where("id = ?", id).Find(&booking).Error
	return
}
