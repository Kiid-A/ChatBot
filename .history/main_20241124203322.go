package main

import (
	"ChatBot/internal/db"
	"ChatBot/internal/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var db db.MyDB
	var ctrl controller.Controller
	db.InitDB()
	ctrl.InitCtrl(db)

	r := gin.Default()
	r.Use(cors.Default())

	// r.LoadHTMLGlob("web/src/*.html")
	// r.Static("assets/", "web/src/assets/")
	controller.LoadIndex(r)
	ctrl.Uc.LoadUser(r)

	r.Run(":9998")
}
