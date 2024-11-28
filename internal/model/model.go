package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string `gorm:"column:id;" json:"id"`
	Name     string `gorm:"column:name;" json:"name"`
	Username string `gorm:"column:username; primary_key;" json:"username"`
	Passwd   string `gorm:"column:passwd;" json:"password"`

	Balance float64 `gorm:"column:balance;" json:"balance"`
}

type ChatMessage struct {
	gorm.Model
	UserId  string `gorm:"column:user_id;" json:"user_id"`
	ChatId  string `gorm:"column:chat_id;" json:"chat_id"`
	Content string `gorm:"column:content;" json:"content"`
}
