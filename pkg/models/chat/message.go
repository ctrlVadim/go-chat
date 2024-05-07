package models

type Message struct {
	Id        int    `json:"id"`
	ChatId    int    `json:"chat_id" binding:"required"`
	Content   string `json:"message" binding:"required"`
	CreatedAt int64  `json:"created_at"`
}
