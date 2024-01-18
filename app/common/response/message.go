package response

import "registration-booking/app/models"

type ConversationsRes struct {
	Conversations []Conversations `json:"conversations"`
	Count         int64           `json:"count"`
}

type Conversations struct {
	Conversation models.Message `json:"conversation"`
	UnreadCount  int64          `json:"unread_count"`
	LetterCount  int64          `json:"letter_count"`
	Target       models.User    `json:"target"`
}

type UnreadCountRes struct {
	LetterUnreadCount int64 `json:"letter_unread_count"`
	NoticeUnreadCount int64 `json:"notice_unread_count"`
}
