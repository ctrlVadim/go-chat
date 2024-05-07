package models

type Chat struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}
