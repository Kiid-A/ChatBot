package service

import (
	"github.com/gin-gonic/gin"
	"ChatBot/api/"
)


type UserService interface {
	Register(ctx gin.Context, req *RegisterRequest)
}
