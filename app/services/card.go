package services

import (
	"registration-booking/app/common/request"
	"registration-booking/app/models"
	"registration-booking/global"
	"strconv"
)

type cardService struct{}

var CardService = new(cardService)

// CreateCard 创建就诊卡
func (s *cardService) CreateCard(form request.Card, id string) (error, models.Card) {
	userId, _ := strconv.Atoi(id)
	card := models.Card{
		UserId:       userId,
		Name:         form.Name,
		IdType:       form.IdType,
		IdNumber:     form.IdNumber,
		Nation:       form.Nation,
		Relationship: form.Relationship,
		Phone:        form.Phone,
		Address:      form.Address,
	}
	if err := global.App.DB.Create(&card).Error; err != nil {
		return err, card
	}
	return nil, card
}

// GetCardList 获取就诊卡列表
func (s *cardService) GetCardList(id string) (error, []models.Card) {
	var list []models.Card
	if err := global.App.DB.Where("user_id = ?", id).Find(&list).Error; err != nil {
		return err, list
	}
	return nil, list
}

// GetCardById 获取就诊卡详情
func (s *cardService) GetCardById(id string) (error, models.Card) {
	var card models.Card
	if err := global.App.DB.First(&card, id).Error; err != nil {
		return err, card
	}
	return nil, card
}

// UpdateCard 更新就诊卡
func (s *cardService) UpdateCard(card models.Card, form request.Card) error {
	if err := global.App.DB.Model(&card).Updates(form).Error; err != nil {
		return err
	}
	return nil
}

// DeleteCard 删除就诊卡
func (s *cardService) DeleteCard(id string) error {
	if err := global.App.DB.Delete(&models.Card{}, id).Error; err != nil {
		return err
	}
	return nil
}
