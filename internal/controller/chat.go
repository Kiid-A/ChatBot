package controller

import (
	"ChatBot/cmd"
	"ChatBot/internal/model"
	"strconv"

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
	e.GET("/:id/history", cc.History)
}

func (cc *ChatController) NewChat(ctx *gin.Context) {
	var chatId string
	userId := ctx.Param("id")

	req := model.ChatRequest{}
	req.UserId = userId

	var cnt int64
	if err := cc.DB.
		Table("chat_messages").
		Model(&model.ChatMessage{}).
		Where("user_id=?", req.UserId).
		Select("MAX(chat_id)").Scan(&cnt).Error; err != nil {
		cnt = 0
	}
	chatId = strconv.Itoa(int(cnt + 1))

	// call python
	if err := cmd.RunBot(); err != nil {
		ctx.JSON(500, gin.H{"error": "failed to run chatbot" + err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"chat_id": chatId,
	})
}

func (cc *ChatController) EndChat(ctx *gin.Context) {
	// userId := ctx.Param("id")
	// chatId := ctx.Param("chat_id")

	// close chat in python
	if err := cmd.EndBot(); err != nil {
		ctx.JSON(500, gin.H{"error": "failed to close chatbot" + err.Error()})
		return
	}

	ctx.JSON(200, "End Chat Succesfully!")
}

func (cc *ChatController) Chat(ctx *gin.Context) {
	var reply model.ChatReply
	userId := ctx.Param("id")
	chatId := ctx.Param("chat_id")
	req := model.ChatRequest{}
	ctx.BindJSON(&req)
	req.UserId = userId
	req.ChatId = chatId

	mes := model.ChatMessage{
		ChatId:  chatId,
		UserId:  userId,
		Content: req.Question,
	}
	res := cc.DB.Table("chat_messages").Create(&mes)
	if res.Error != nil {
		panic(res.Error)
	}

	// chat
	reply, err := cmd.AskBot(req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to ask chatbot" + err.Error()})
		return
	}

	// save reply
	mes = model.ChatMessage{
		ChatId:  chatId,
		UserId:  userId,
		Content: reply.Response,
	}
	res = cc.DB.Table("chat_messages").Create(&mes)
	if res.Error != nil {
		panic(res.Error)
	}

	ctx.JSON(200, gin.H{
		"answer": reply.Response,
	})
}

func (cc *ChatController) History(ctx *gin.Context) {
	userId := ctx.Param("id")

	records := []model.ChatMessage{}

	cc.DB.
		Table("chat_messages").
		Where("user_id=?", userId).
		Order("created_at ASC").
		Find(&records)

	ctx.JSON(200, records)
}
