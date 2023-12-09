package models

import (
	"registration-booking/global"
	"strconv"
)

type Article struct {
	ID
	Title   string `json:"title" gorm:"index;not null;comment:标题"`
	URL     string `json:"url" gorm:"not null;comment:文章地址"`
	Photo   string `json:"photo" gorm:"not null;comment:图片地址"`
	Content string `json:"content" gorm:"not null;comment:内容"`
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
