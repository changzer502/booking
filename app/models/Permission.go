package models

import "strconv"

type Permission struct {
	ID
	PermissionName string `json:"permission_name" gorm:"not null;comment:权限名称"`
	Sort           int    `json:"sort" gorm:"not null;comment:排序"`
	ParentId       int    `json:"parent_id" gorm:"not null;comment:父级ID"`
	Icon           string `json:"icon" gorm:"comment:图标"`
	Router         string `json:"router" gorm:"not null;comment:路由"`
	Timestamps
	SoftDeletes
}

func (user Permission) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
