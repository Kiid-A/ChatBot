package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChatController struct {
	DB *gorm.DB
}

func (cc *ChatController) LoadChat(e *gin.Engine) {
	e.GET("/chat.html", func(ctx *gin.Context) {
		ctx.HTML(200, "chat.html", gin.H{})
	})
	e.GET("/newchat.html", func(ctx *gin.Context) {
		ctx.HTML(200, "newchat.html", gin.H{})
	})
	e.GET("/endchat.html", func(ctx *gin.Context) {
		ctx.HTML(200, "endchat.html", gin.H{})
	})

	e.POST("/:id/newchat", cc.NewChat)
	e.POST("/:id/:chat_id/chat", cc.Chat)
	e.POST("/:id/:chat_id/endchat", cc.EndChat)
}

func (cc *ChatController) NewChat(ctx *gin.Context) {
	var chatId string
	userId := ctx.Param("id")

	// call python

	ctx.JSON(200, chatId)
}

func (cc *ChatController) EndChat(ctx *gin.Context) {
	userId := ctx.Param("id")
	chatId := ctx.Param("")
	// close chat in python
	
	ctx.JSON(200, "End Chat Succesfully!")	
}

func (cc *ChatController) Chat(ctx *gin.Context) {
	var ans string

}
