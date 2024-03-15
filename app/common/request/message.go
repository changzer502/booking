package request

type SendMessageReq struct {
	ToUserId string `json:"to_user_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type SendNoticeReq struct {
	ToUserId string `json:"to_user_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
}
