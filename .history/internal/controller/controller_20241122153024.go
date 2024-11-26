package controller

import (
	"ChatBot/internal/db"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Uc UserController
}

func (c *Controller) InitCtrl(db db.MyDB) {
	c.Ac = ArticleController{
		DB: db.ArticleDB,
	}
	c.Uc = UserController{
		DB: db.UserDB,
	}
	c.Mc = MapController{
		DB: db.MapDB,
	}
}

func LoadIndex(e *gin.Engine) {
	e.GET("/index.html", welcome)
	e.GET("/", welcome)
}

func welcome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
