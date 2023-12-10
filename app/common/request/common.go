package request

type Page struct {
	Page      int `form:"page" json:"page"`
	PageNo    int `form:"pageNo" json:"pageNo"`
	PageSize  int `form:"page_size" json:"page_size"`
	PageSize2 int `form:"pageSize" json:"pageSize" `
}

func (page Page) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"page_size.required": "每页数量不能为空",
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
