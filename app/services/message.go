package services

import (
	"errors"
	"fmt"
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

func (service *messageService) UnreadCount(uid string) (*response.UnreadCountRes, error) {
	//查询未读信息数量
	letterUnreadCount, err := models.FindLetterUnreadCount(uid, "")
	if err != nil {
		return nil, err
	}
	noticeUnreadCount, err := models.FindNoticeUnreadCount(uid, "notice")
	if err != nil {
		return nil, err
	}
	return &response.UnreadCountRes{
		LetterUnreadCount: letterUnreadCount,
		NoticeUnreadCount: noticeUnreadCount,
	}, nil
}

func (service *messageService) GetConversationDetail(uid, conversationId string, pageNo, pageSize int) (*response.ConversationsRes, error) {
	userId, _ := strconv.Atoi(uid)
	letters, count, err := models.FindLetters(conversationId, pageNo, pageSize)
	if err != nil {
		return nil, err
	}
	conversations := make([]response.Conversations, 0)
	unreadIds := make([]uint, 0)
	for _, letter := range letters {
		conversation := response.Conversations{
			Conversation: letter,
		}
		// 对方用户
		targetId := letter.FromId
		if uint(userId) == letter.FromId {
			targetId = letter.ToId
		}
		target, err := models.FindUserById(targetId)
		if err != nil {
			return nil, err
		}
		conversation.Target = *target

		conversations = append(conversations, conversation)

		//获得未读消息的id
		if uint(userId) == letter.ToId && letter.Status == "0" {
			unreadIds = append(unreadIds, letter.ID.ID)
		}
	}

	// 更新未读消息状态
	if len(unreadIds) > 0 {
		err = models.UpdateMessageStatus(unreadIds, models.Read)
		if err != nil {
			return nil, err
		}
	}
	// 未读数量
	unreadCount, err := models.FindLetterUnreadCount(uid, conversationId)
	if err != nil {
		return nil, err
	}
	return &response.ConversationsRes{
		Count:         count,
		Conversations: conversations,
		UnReadCount:   unreadCount,
	}, nil
}

func (service *messageService) SendLetter(uid, toUserId, content string) error {
	toId, _ := strconv.Atoi(toUserId)
	userId, _ := strconv.Atoi(uid)
	toUser, err := models.FindUserById(uint(toId))
	if err != nil {
		return errors.New(fmt.Sprintf("用户不存在"))
	}
	// 创建消息
	message := models.Message{
		FromId:         uint(userId),
		ToId:           toUser.ID.ID,
		Content:        content,
		Status:         "0",
		ConversationId: getConversationId(toId, userId),
	}
	err = message.CreateMessage()
	if err != nil {
		return errors.New("创建消息失败")
	}
	return nil
}

func (service *messageService) SendNotice(uid, toUserId, content string) error {
	toId, _ := strconv.Atoi(toUserId)
	userId, _ := strconv.Atoi(uid)
	toUser, err := models.FindUserById(uint(toId))
	if err != nil {
		return errors.New(fmt.Sprintf("用户不存在"))
	}
	// 创建消息
	message := models.Message{
		FromId:         uint(userId),
		ToId:           toUser.ID.ID,
		Content:        content,
		Status:         "0",
		ConversationId: "notice",
	}
	err = message.CreateMessage()
	if err != nil {
		return errors.New("创建消息失败")
	}
	return nil
}

func getConversationId(id1, id2 int) string {
	if id1 > id2 {
		return fmt.Sprintf("%v_%v", id2, id1)
	}
	return fmt.Sprintf("%v_%v", id1, id2)

}

func (service *messageService) GetNoticeList(uid string, pageNo, pageSize int) ([]models.Message, int64, error) {
	//会话列表
	conversationList, count, err := models.FindNotices(uid, pageNo, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return conversationList, count, nil
}
