package handler

import (
	"github.com/gin-gonic/gin"
	"registration-booking/app/common/request"
	"registration-booking/app/common/response"
	"registration-booking/app/services"
)

// 获得科室列表
func GetDepartmentList(c *gin.Context) {
	if err, list := services.DepartmentService.GetDepartmentList(); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, list)
	}
}

// 创建科室
func CreateDepartment(c *gin.Context) {
	var form request.Department
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, department := services.DepartmentService.CreateDepartment(form, c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, department)
	}
}

func GetDepartmentPage(c *gin.Context) {
	var form request.Page
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, department := services.DepartmentService.GetDepartmentPage(form); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, department)
	}
}

func GetDepartmentById(c *gin.Context) {
	if err, department := services.DepartmentService.GetDepartmentById(c.Param("id")); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, department)
	}
}
