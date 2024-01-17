package services

import (
	"registration-booking/app/common/response"
	"registration-booking/app/models"
	"strconv"
)

type messageService struct {
}

var MessageService = new(messageService)

func (service *messageService) GetLetterList(uid string, pageNo, pageSize int) (*response.ConversationsRes, error) {
	//会话列表
	conversationList, count, err := models.FindConversations(uid, pageNo, pageSize)
	if err != nil {
		return nil, err
	}
	conversations := make([]response.Conversations, 0)
	for _, message := range conversationList {
		conversation := response.Conversations{
			Conversation: message,
		}
		// 未读数量
		unreadCount, err := models.FindLetterUnreadCount(uid, message.ConversationId)
		if err != nil {
			return nil, err
		}
		conversation.UnreadCount = unreadCount
		// 全部数量
		letterCount, err := models.FindLetterCount(message.ConversationId)
		if err != nil {
			return nil, err
		}
		conversation.LetterCount = letterCount
		// 对方用户
		targetId := message.FromId
		userId, _ := strconv.Atoi(uid)
		if uint(userId) == message.FromId {
			targetId = message.ToId
		}
		target, err := models.FindUserById(targetId)
		if err != nil {
			return nil, err
		}
		conversation.Target = *target

		conversations = append(conversations, conversation)
	}
	return &response.ConversationsRes{
		Count:         count,
		Conversations: conversations,
	}, nil
}
