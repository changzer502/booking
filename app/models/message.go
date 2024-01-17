package models

import (
	"strconv"
)

type Message struct {
	ID
	FromId         uint   `json:"from_id" gorm:"index:comment:发送者ID"`
	ToId           uint   `json:"to_id" gorm:"index:comment:接收者ID"`
	ConversationId uint   `json:"conversation_id" gorm:"index:comment:会话ID"`
	Content        string `json:"content" gorm:"comment:消息内容"`
	Status         string `json:"status" gorm:"comment:消息状态（0未读，1已读）" default:"0" `
	Timestamps
	SoftDeletes
}

func (message Message) GetUid() string {
	return strconv.Itoa(int(message.ID.ID))
}
