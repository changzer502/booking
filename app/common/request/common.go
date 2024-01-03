package request

type Page struct {
	Query    string `json:"query"`
	PageNo   int    `form:"pageNo" json:"pageNo" binding:"required"`
	PageSize int    `form:"pageSize" json:"pageSize"  binding:"required"`
}

func (page Page) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"pageNo.required":   "第几页不能为空",
		"pageSize.required": "每页数量不能为空",
	}
}

type Id struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

func (id Id) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"id.required": "id 不能为空",
	}
}
