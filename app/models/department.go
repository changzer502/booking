package models

import (
	"registration-booking/global"
	"strconv"
)

type Department struct {
	ID
	DeptName     string       `json:"dept_name" gorm:"not null;comment:部门名称"`
	Icon         string       `json:"icon" gorm:"not null;comment:部门图标"`
	ParentId     uint         `json:"parent_id" gorm:"index;not null;comment:上级部门id"`
	OrderNum     int          `json:"order_num" gorm:"index;not null;comment:排序"`
	Ancestors    string       `json:"ancestors" gorm:"not null;comment:祖级列表"`
	Introduction string       `json:"introduction" gorm:"not null;comment:部门介绍"`
	Children     []Department `json:"children" gorm:"-"` // 忽略本字段
	Timestamps
	SoftDeletes
}

func (department Department) GetUid() string {
	return strconv.Itoa(int(department.ID.ID))
}
func FindChildrenByDeptId(id uint) (department []Department, err error) {
	err = global.App.DB.Where("parent_id = ?", id).Find(&department).Error
	return
}
func FindDepartmentById(id uint) (department Department, err error) {
	err = global.App.DB.Where("id = ?", id).Find(&department).Error
	return
}
func UpdateDepartment(department Department) (err error) {
	err = global.App.DB.Model(&department).Updates(&department).Error
	return
}
func DeleteDepartment(department Department) (err error) {
	err = global.App.DB.Delete(&department).Error
	return
}
