package request

type Card struct {
	Id           int    `form:"id" json:"id"`
	Name         string `form:"name" json:"name" binding:"required"`
	IdType       string `form:"id_type" json:"id_type" binding:"required"`
	IdNumber     string `form:"id_number" json:"id_number" binding:"required"`
	Nation       string `form:"nation" json:"nation" binding:"required"`
	Relationship int    `form:"relationship" json:"relationship" `
	Phone        string `form:"phone" json:"phone" binding:"required"`
	Address      string `form:"address" json:"address" binding:"required"`
	Default      bool   `form:"default" json:"default"`
}

func (login Card) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":      "姓名不能为空",
		"id_type.required":   "证件类型不能为空",
		"id_number.required": "证件号码不能为空",
		"nation.required":    "民族不能为空",
		"phone.required":     "联系电话不能为空",
		"address.required":   "联系地址不能为空",
	}
}
