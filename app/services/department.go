package services

import (
	"registration-booking/app/common/request"
	"registration-booking/app/common/response"
	"registration-booking/app/models"
	"registration-booking/global"
	"strconv"
	"time"
)

type departmentService struct{}

var DepartmentService = new(departmentService)

// GetDepartmentList 获取科室列表
func (s *departmentService) GetDepartmentList() (error, []models.Department) {
	var list []models.Department
	// 获得第一级科室
	if err := global.App.DB.Where("parent_id = ?", 0).Find(&list).Error; err != nil {
		return err, list
	}
	// 获得第二级科室
	for i := 0; i < len(list); i++ {
		var children []models.Department
		if err := global.App.DB.Where("parent_id = ?", list[i].ID.ID).Find(&children).Error; err != nil {
			return err, list
		}
		list[i].Children = children
	}
	return nil, list
}

func (s *departmentService) CreateDepartment(department request.Department, id string) (error, models.Department) {
	userId, _ := strconv.Atoi(id)
	dept := models.Department{
		DeptName:     department.DeptName,
		Icon:         department.Icon,
		ParentId:     department.ParentId,
		Introduction: department.Introduction,
		OrderNum:     0,
		Timestamps: models.Timestamps{
			CreatedAt: time.Now(),
			CreatedBy: uint(userId),
			UpdatedAt: time.Now(),
			UpdatedBy: uint(userId),
		},
	}
	if dept.ParentId == 0 {
		dept.Ancestors = "0"
	} else {
		var parent models.Department
		if err := global.App.DB.First(&parent, dept.ParentId).Error; err != nil {
			return err, dept
		}
		dept.Ancestors = parent.Ancestors + "," + parent.GetUid()
	}
	if err := global.App.DB.Create(&dept).Error; err != nil {
		return err, dept
	}
	return nil, dept
}

func (s *departmentService) GetDepartmentPage(page request.Page) (error, *response.PageData) {
	var list []models.Department
	if err := global.App.DB.Where("parent_id != 0").Order("LENGTH(dept_name)").Offset((page.Page - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error; err != nil {
		return err, nil
	}
	var total int64
	if err := global.App.DB.Model(&models.Department{}).Count(&total).Error; err != nil {
		return err, nil
	}

	return nil, &response.PageData{
		List:     list,
		Total:    total,
		PageData: list,
	}
}

func (s *departmentService) GetDepartmentById(id string) (error, *models.Department) {
	dept := new(models.Department)
	if err := global.App.DB.Where("id = " + id).Find(&dept).Error; err != nil {
		return err, nil
	}
	return nil, dept
}
