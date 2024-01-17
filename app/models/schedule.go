package models

import (
	"registration-booking/global"
	"strconv"
)

type Schedule struct {
	ID
	Week         int    `json:"week" gorm:"not null;comment:星期"`
	DoctorId     uint   `json:"doctor_id" gorm:"not null;comment:医生Id"`
	DepartmentId uint   `json:"department_id" gorm:"not null;comment:科室Id"`
	Price        int    `json:"price" gorm:"default:10;comment:价格"`
	Time         string `json:"time" gorm:"not null;comment:时间（上午/下午）"`
	Timestamps
	SoftDeletes
}

func (schedule Schedule) GetUid() string {
	return strconv.Itoa(int(schedule.ID.ID))
}

func FindSchedulesByDoctorId(doctorId, day string) (ScheduleList []Schedule, err error) {
	err = global.App.DB.Where("doctor_id = ? AND day = ?", doctorId, day).Find(&ScheduleList).Error
	return
}
func FindAllSchedulesByDoctorId(doctorId string) (ScheduleList []Schedule, err error) {
	err = global.App.DB.Where("doctor_id = ?", doctorId).Order("time").Find(&ScheduleList).Error
	return
}
func FindSchedulesByDepartmentID(departmentId uint, week int) (ScheduleList []Schedule, err error) {
	err = global.App.DB.Where("department_id = ? AND week = ?", departmentId, week).Find(&ScheduleList).Error
	return
}
func FindScheduleByID(id uint) (schedule Schedule, err error) {
	err = global.App.DB.Where("id = ?", id).Find(&schedule).Error
	return
}
