package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string `gorm:"column:id;" json:"id"`
	Name     string `gorm:"column:name;" json:"name"`
	Username string `gorm:"column:username; primary_key;" json:"username"`
	Passwd   string `gorm:"column:passwd;" json:"password"`
}

type NewChatReq struct {

}

type EndChatReq struct {
	
}
