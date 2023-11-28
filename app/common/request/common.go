package request

type Page struct {
	Page     int `form:"page" json:"page" binding:"required"`
	PageSize int `form:"page_size" json:"page_size" binding:"required"`
}

func (page Page) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"page.required":      "页码不能为空",
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
