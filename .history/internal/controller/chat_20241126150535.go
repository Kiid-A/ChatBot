package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChatController struct {
	DB *gorm.DB
}

func (cc *ChatController) LoadChat(e *gin.Engine) {
	e.POST("/:id/newchat", cc.NewChat)
	e.POST("/:id/:chat_id/chat", cc.Chat)
	e.POST("/:id/:chat_id/endchat", cc.)
}

func (cc *ChatController) NewChat(ctx *gin.Context) {

}

func (cc *ChatController) EndChat(ctx *gin.Context) {
	
}

func (cc *ChatController) Chat(ctx *gin.Context) {
	
}
