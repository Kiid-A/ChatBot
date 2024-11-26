package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChatController struct {
	DB *gorm.DB
}

func (c *ChatController) LoadChat(e *gin.Engine) {
	e.POST("/:id/newchat", c.NewChat)
}

func (cc *ChatController) NewChat(c *gin.Context) {

}

func (c *ChatController) EndChat(e *gin.Engine) {
	
}

func (c *ChatController) Chat(e *gin.Engine) {
	
}
