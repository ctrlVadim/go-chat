package models

type Participant struct {
	Id     int `json:"id"`
	UserId int `json:"user_id" binding:"required"`
	ChatId int `json:"chat_id" binding:"required"`
}
