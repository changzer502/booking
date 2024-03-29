package models

import (
	"registration-booking/global"
	"strconv"
)

type Message struct {
	ID
	FromId         uint   `json:"from_id" gorm:"index:comment:发送者ID"`
	ToId           uint   `json:"to_id" gorm:"index:comment:接收者ID"`
	ConversationId string `json:"conversation_id" gorm:"index:comment:会话ID"`
	Content        string `json:"content" gorm:"comment:消息内容"`
	Status         string `json:"status" gorm:"comment:消息状态（0未读，1已读）" default:"0" `
	Timestamps
	SoftDeletes
}

type Content struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	BookingID int64  `json:"bookingId"`
}

const (
	UnRead = iota
	Read
)

func (message Message) GetUid() string {
	return strconv.Itoa(int(message.ID.ID))
}
func (message Message) CreateMessage() (err error) {
	err = global.App.DB.Create(&message).Error
	return
}

func FindConversations(uid string, page, pageSize int) (messages []Message, count int64, err error) {
	err = global.App.DB.Where("id in (select max(id) from messages where status != 2 and from_id != 1  and (from_id = ? or to_id = ?) group by conversation_id) ", uid, uid).Order("messages.id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&messages).Count(&count).Error
	return
}

func FindLetterUnreadCount(uid, countonversationId string) (count int64, err error) {
	query := ""
	if countonversationId != "" {
		query += "and conversation_id = '" + countonversationId + "'"
	}
	err = global.App.DB.Table("messages").Where("status = 0 and from_id != 1 and to_id = ? "+query, uid).Count(&count).Error
	return
}
func FindLetterCount(countonversationId string) (count int64, err error) {
	err = global.App.DB.Table("messages").Where("status != 2 and from_id != 1 and conversation_id = ?", countonversationId).Count(&count).Error
	return
}

func FindNoticeUnreadCount(uid, topic string) (count int64, err error) {
	query := ""
	if topic != "" {
		query += "and conversation_id = '" + topic + "'"
	}
	err = global.App.DB.Table("messages").Where("status = 0 and from_id = 1 and to_id = ? "+query, uid).Count(&count).Error
	return
}
func FindLetters(countonversationId string, page, pageSize int) (message []Message, count int64, err error) {
	err = global.App.DB.Where("status != 2 and from_id != 1 and conversation_id = ?", countonversationId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&message).Count(&count).Error
	return
}

func UpdateMessageStatus(ids []uint, status int) (err error) {
	err = global.App.DB.Table("messages").Where("id in (?)", ids).Update("status", status).Error
	return
}

func FindNotices(uid string, page, pageSize int) (messages []Message, count int64, err error) {
	err = global.App.DB.Table("messages").Where("from_id = 1 and conversation_id = 'notice' and to_id = ?", uid).Offset((page - 1) * pageSize).Limit(pageSize).Find(&messages).Count(&count).Error
	return
}
