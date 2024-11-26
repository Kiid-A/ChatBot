package main

import (
	"ChatBot/config"
	"ChatBot/internal/controller"
	"ChatBot/internal/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 应用初始化
	var db db.MyDB
	var ctrl controller.Controller
	db.InitDB()
	ctrl.InitCtrl(db)
	
	r := gin.Default()
	r.Use(cors.Default())

	r.LoadHTMLGlob("web/src/*.html")
	r.Static("assets/", "web/src/assets/")
	controller.LoadIndex(r)
	ctrl.Uc.LoadUser(r)

	r.Run(":" + config.ServerPort)
}
