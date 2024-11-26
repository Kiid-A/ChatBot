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

	// 挂载跨域
	r := gin.Default()
	r.Use(cors.Default())

	// 挂载资源路径
	r.LoadHTMLGlob("web/src/*.html")
	r.Static("assets/", "web/src/assets/")
	ctrl.LoadAll(r)

	r.Run(":" + config.ServerPort)
}
