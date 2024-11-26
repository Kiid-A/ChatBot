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

func (c *ChatController)
