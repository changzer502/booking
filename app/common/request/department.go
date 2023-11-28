package request

type Department struct {
	DeptName     string `form:"dept_name" json:"dept_name" binding:"required" `
	Icon         string `form:"icon" json:"icon" binding:"required" `
	ParentId     uint   `form:"parent_id" json:"parent_id"  `
	Introduction string `form:"introduction" json:"introduction" binding:"required" `
}

func (department Department) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"dept_name.required":    "科室名不能为空",
		"icon.required":         "科室图标不能为空",
		"introduction.required": "科室介绍不能为空",
	}
}
