package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChatController struct {
	DB *gorm.DB
}

func (c *ChatController) LoadChat(e *gin.Engine) {
		
}

func (c *ChatController) NewChat(e *gin.Engine) {

}

func (c *ChatController) EndChat(e *gin.Engine) {
	
}

func (c *ChatController) Chat(e *gin.Engine) {
	
}
