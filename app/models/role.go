package models

import "strconv"

type Role struct {
	ID
	RoleName    string `json:"role_name" gorm:"not null;comment:角色名称"`
	Description string `json:"description" gorm:"comment:角色描述"`
	Timestamps
	SoftDeletes
}

func (user Role) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
