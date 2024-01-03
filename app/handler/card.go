package handler

import (
	"github.com/gin-gonic/gin"
	"registration-booking/app/common/request"
	"registration-booking/app/common/response"
	"registration-booking/app/models"
	"registration-booking/app/services"
)

// CreateCard 创建就诊卡
func CreateCard(c *gin.Context) {
	var form request.Card
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.CardService.CreateCard(form, c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

// GetCardList 获取就诊卡列表
func GetCardList(c *gin.Context) {

	if err, list := services.CardService.GetCardList(c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, list)
	}
}

func GetAllCardList(c *gin.Context) {
	var form request.Page
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, total, list := services.CardService.GetAllCardList(form); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, response.PageData{PageData: list, Total: total})
	}
}

// GetCardById 获取就诊卡详情
func GetCardById(c *gin.Context) {
	if err, card := services.CardService.GetCardById(c.Param("id")); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, card)
	}
}

// UpdateCard 更新就诊卡
func UpdateCard(c *gin.Context) {
	var form request.Card
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, card := services.CardService.GetCardById(c.Param("id")); err != nil {
		response.Fail(c, err.Error())
	} else {
		if form.Default {
			// 取消其他默认就诊卡
			err := models.CancelOtherCardDefaults(uint(card.UserId), card.ID.ID)
			if err != nil {
				return
			}
		} else if card.Default {
			err := models.CancelCardDefault(uint(card.UserId), card.ID.ID)
			if err != nil {
				return
			}
		}
		if err := services.CardService.UpdateCard(card, form); err != nil {
			response.Fail(c, err.Error())
		} else {
			response.Success(c, nil)
		}
	}
}

// DeleteCard 删除就诊卡
func DeleteCard(c *gin.Context) {
	if err := services.CardService.DeleteCard(c.Param("id")); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, nil)
	}
}
