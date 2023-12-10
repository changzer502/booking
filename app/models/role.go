package models

import (
	"registration-booking/global"
	"strconv"
)

type Role struct {
	ID
	RoleName    string `json:"role_name" gorm:"not null;comment:角色名称"`
	Description string `json:"description" gorm:"comment:角色描述"`
	RoleKey     string `json:"role_key" gorm:"comment:角色标识"`
	Timestamps
	SoftDeletes
}

func (user Role) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}

func GetRoleById(id uint) (role Role, err error) {
	err = global.App.DB.Where("id = ?", id).First(&role).Error
	return
}
