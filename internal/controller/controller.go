package controller

import (
	"ChatBot/internal/db"

	"github.com/gin-gonic/gin"
)

// Main Controller
// 管理Controller, 负责各类Controller的初始化

type Controller struct {
	U UserController
	C ChatController
}

func (c *Controller) InitCtrl(db db.MyDB) {
	c.U = UserController{
		DB: db.UserDB,
	}
	c.C = ChatController{
		DB: db.ChatDB,
	}
}

func (c *Controller) LoadAll(e *gin.Engine) {
	LoadIndex(e)
	c.U.LoadUser(e)
	c.C.LoadChat(e)
}

func LoadIndex(e *gin.Engine) {
	e.GET("/index.html", welcome)
	e.GET("/", welcome)
}

func welcome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
