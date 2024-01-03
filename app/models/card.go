package models

import (
	"registration-booking/global"
	"strconv"
)

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
	Default      bool   `json:"default" gorm:"not null;comment:是否默认"`
	Timestamps
	SoftDeletes
}

func (user Card) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
func FindCardById(id uint) (card Card, err error) {
	err = global.App.DB.Where("id = ?", id).Find(&card).Error
	return
}
func CancelAllCardDefaults(uid uint) error {
	return global.App.DB.Model(&Card{}).Where("user_id = ?", uid).Update("default", false).Error
}

func CancelOtherCardDefaults(uid, cardId uint) error {
	return global.App.DB.Model(&Card{}).Where("user_id = ? AND id != ?", uid, cardId).Update("default", false).Error
}

func CancelCardDefault(uid, cardId uint) error {
	return global.App.DB.Model(&Card{}).Where("user_id = ? AND id = ?", uid, cardId).Update("default", false).Error
}

func FindAllCardListByPage(page, pageSize int, query string) (cards []Card, total int64, err error) {
	err = global.App.DB.Where("1 = 1" + query).Offset((page - 1) * pageSize).Limit(pageSize).Find(&cards).Count(&total).Error
	return
}
