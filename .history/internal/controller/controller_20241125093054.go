package controller

import (
	"ChatBot/internal/db"

	"github.com/gin-gonic/gin"
)

// Main Controller
// 管理Controller, 负责各类Controller的初始化

type Controller struct {
	Uc UserController
}

func (c *Controller) InitCtrl(db db.MyDB) {
	c.Uc = UserController{
		DB: db.UserDB,
	}
}

func (c *Controller) LoadAll(e *gin.Engine) {
	LoadIndex(e)
	c.Uc.LoadUser(e)
	c.C
}

func LoadIndex(e *gin.Engine) {
	e.GET("/index.html", welcome)
	e.GET("/", welcome)
}

func welcome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
