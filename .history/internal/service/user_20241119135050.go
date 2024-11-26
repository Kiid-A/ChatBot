package service

import (
	"github.com/gin-gonic/gin"
	"Chat"
)


type UserService interface {
	Register(ctx gin.Context, req *RegisterRequest)
}
