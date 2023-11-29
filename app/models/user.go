package models

import (
	"registration-booking/global"
	"strconv"
)

type User struct {
	ID
	Nickname     string `json:"nickname" gorm:"not null;comment:用户昵称"`
	Gender       int    `json:"gender" gorm:"comment:用户性别"`
	AvatarUrl    string `json:"avatar_url" gorm:"not null;comment:用户头像"`
	Mobile       string `json:"mobile" gorm:"comment:用户手机号"`
	OpenId       string `json:"open_id" gorm:"index;comment:WX用户唯一标识"`
	Password     string `json:"password" gorm:"comment:用户密码"`
	RoleId       uint   `json:"role_id" gorm:"comment:用户角色"`
	Introduce    string `json:"introduce" gorm:"size:2048;comment:用户简介"`
	DepartmentID uint   `json:"department_id" gorm:"index;comment:医生所属科室"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}

func FindDoctorsByDepartmentId(departmentId string) (doctorList []User, err error) {
	err = global.App.DB.Where("department_id = ? AND role_id = ?", departmentId, 2).Find(&doctorList).Error
	return
}

func FindDoctorById(id uint) (doctor User, err error) {
	err = global.App.DB.Where("id = ? AND role_id = ?", id, 2).Find(&doctor).Error
	return
}
