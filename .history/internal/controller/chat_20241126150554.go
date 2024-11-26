package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChatController struct {
	DB *gorm.DB
}

func (cc *ChatController) LoadChat(e *gin.Engine) {
	e.GET("/login.html", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", gin.H{})
	})
	e.GET("/register.html", func(ctx *gin.Context) {
		ctx.HTML(200, "register.html", gin.H{})
	})
	e.GET("/user_info.html", func(ctx *gin.Context) {
		ctx.HTML(200, "user_info.html", gin.H{})
	})

	e.POST("/:id/newchat", cc.NewChat)
	e.POST("/:id/:chat_id/chat", cc.Chat)
	e.POST("/:id/:chat_id/endchat", cc.EndChat)
}

func (cc *ChatController) NewChat(ctx *gin.Context) {

}

func (cc *ChatController) EndChat(ctx *gin.Context) {
	
}

func (cc *ChatController) Chat(ctx *gin.Context) {
	
}
