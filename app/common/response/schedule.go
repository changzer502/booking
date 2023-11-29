package response

import "registration-booking/app/models"

type ScheduleList struct {
	Doctor                models.User         `json:"doctor"`
	ScheduleAndTicketList []ScheduleAndTicket `json:"schedule_and_ticket_list"`
}
type ScheduleAndTicket struct {
	Schedule models.Schedule `json:"schedule"`
	Ticket   models.Ticket   `json:"ticket"`
	Status   bool            `json:"status"`
}
