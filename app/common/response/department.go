package response

import "registration-booking/app/models"

type DepartmentRes struct {
	models.Department
	ParentName string `json:"parent_name"`
}
