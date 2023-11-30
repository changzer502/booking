package request

type CreateScheduleReq struct {
	DoctorId     uint   `form:"doctor_id" json:"doctor_id" binding:"required"`
	Week         int    `form:"week" json:"week" binding:"required"`
	Time         string `form:"time" json:"time" binding:"required"`
	Price        int    `form:"price" json:"price" binding:"required"`
	DepartmentId uint   `form:"department_id" json:"department_id" binding:"required"`
}

type GetScheduleListReq struct {
	DepartmentId uint   `form:"department_id" json:"department_id" binding:"required"`
	Week         int    `form:"week" json:"week" binding:"required"`
	Day          string `form:"day" json:"day" binding:"required"`
	UserId       uint   `form:"user_id" json:"user_id"`
	CardId       int    `form:"card_id" json:"card_id"`
	Only         bool   `form:"only" json:"only"` // 只看有号的
}

type CreateTicketReq struct {
	ScheduleId uint   `form:"schedule_id" json:"schedule_id" binding:"required"`
	Day        string `form:"day" json:"day" binding:"required"`
	Num        int    `form:"num" json:"num" binding:"required"`
}

type BookingReq struct {
	TicketId uint `form:"ticket_id" json:"ticket_id" binding:"required"`
}
