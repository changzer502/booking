package response

import "registration-booking/app/models"

type ScheduleList struct {
	Doctor                models.User         `json:"doctor"`
	ScheduleAndTicketList []ScheduleAndTicket `json:"schedule_and_ticket_list"`
}
type ScheduleAndTicket struct {
	Schedule     models.Schedule `json:"schedule"`
	Ticket       models.Ticket   `json:"ticket"`
	Status       bool            `json:"status"`
	TicketStatus int             `json:"ticket_status"`
}

type TicketInfo struct {
	Doctor   models.User     `json:"doctor"`
	Schedule models.Schedule `json:"schedule"`
	Ticket   models.Ticket   `json:"ticket"`
}

type BookingHistoryRes struct {
	Count        int64         `json:"count"`
	BookingInfos []BookingInfo `json:"bookingInfos"`
}

type BookingInfo struct {
	Card       models.Card       `json:"card"`
	Doctor     models.User       `json:"doctor"`
	Schedule   models.Schedule   `json:"schedule"`
	Ticket     models.Ticket     `json:"ticket"`
	Department models.Department `json:"department"`
	Booking    models.Booking    `json:"booking"`
}

type ScheduleRes struct {
	Doctor    models.User         `json:"doctor"`
	Schedules []ScheduleAndTicket `json:"schedules"`
}
