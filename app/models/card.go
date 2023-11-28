package models

import "strconv"

type Card struct {
	ID
	UserId       int    `json:"user_id" gorm:"index;not null;comment:用户id"`
	Name         string `json:"name" gorm:"not null;comment:姓名"`
	IdType       string `json:"id_type" gorm:"not null;comment:证件类型"`
	IdNumber     string `json:"id_number" gorm:"not null;comment:证件号码"`
	Nation       string `json:"nation" gorm:"not null;comment:民族"`
	Relationship int    `json:"relationship" gorm:"not null;comment:与用户关系(0本人，1其他)"`
	Phone        string `json:"phone" gorm:"not null;comment:联系电话"`
	Address      string `json:"address" gorm:"not null;comment:联系地址"`
	Timestamps
	SoftDeletes
}

func (user Card) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
