package models

import (
	"registration-booking/global"
	"strconv"
)

type Article struct {
	ID
	Title        string `json:"title" gorm:"index;not null;comment:标题"`
	Url          string `json:"url" gorm:"not null;comment:文章地址"`
	Photo        string `json:"photo" gorm:"not null;comment:图片地址"`
	Content      string `json:"content" gorm:"not null;comment:内容"`
	DepartmentId uint   `json:"department_id" gorm:"not null;comment:科室id"`
	Timestamps
	SoftDeletes
}

func (article Article) GetUid() string {
	return strconv.Itoa(int(article.ID.ID))
}
func FindArticleById(id uint) (article Article, err error) {
	err = global.App.DB.Where("id = ?", id).Find(&article).Error
	return
}
func FindArticleByDepartmentAndPage(id uint, page, pageSize int) (article []Article, count int64, err error) {
	err = global.App.DB.Where("department_id = ?", id).Order("created_at DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&article).Count(&count).Error
	return
}

func FindAllArticleByPage(page, pageSize int, query string, deptIds []int) (article []Article, count int64, err error) {
	if len(deptIds) == 0 {
		err = global.App.DB.Where("1 = 1" + query).Order("created_at DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&article).Count(&count).Error
	} else {
		err = global.App.DB.Where("1 = 1"+query+" AND department_id IN (?)", deptIds).Order("created_at DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&article).Count(&count).Error

	}
	return
}
func CreateArticle(article Article) (err error) {
	err = global.App.DB.Create(&article).Error
	return
}
func UpdateArticle(article Article) (err error) {
	err = global.App.DB.Model(&article).Updates(&article).Error
	return
}
func DeleteArticle(id uint) (err error) {
	err = global.App.DB.Delete(&Article{}, id).Error
	return
}
