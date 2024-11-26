package chatbot

import (
	"ChatBot/internal/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	var db db.MyDB
	var ctrl controllers.Controller
	db.InitDB()
	ctrl.InitCtrl(db)

	r := gin.Default()
	r.Use(cors.Default())

	// r.Use()
	r.LoadHTMLGlob("data/src/*.html")
	r.Static("assets/", "data/src/assets/")
	// r.Static("user/assets/", "data/src/assets/")
	controllers.LoadIndex(r)
	ctrl.Uc.LoadUser(r)
	ctrl.Ac.LoadArticle(r)
	ctrl.Mc.LoadMap(r)

	r.Run(":9998")
}
