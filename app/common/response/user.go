package response

import (
	"registration-booking/app/common/request"
	"registration-booking/app/models"
)

type UserRes struct {
	ID           uint     `json:"id"`
	Nickname     string   `json:"nickname"`
	AvatarUrl    string   `json:"avatar_url"`
	DepartmentId uint     `json:"department_id"`
	Role         []string `json:"role"`
}

type DoctorRes struct {
	models.User
	DepartmentName string `json:"dept_name"`
	request.Introduce
}
