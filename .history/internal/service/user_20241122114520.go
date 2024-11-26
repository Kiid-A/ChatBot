package service

import (
	"ChatBot/api"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(ctx gin.Context, req *api.RegisterRequest)
	Login(ctx gin.Context, req *api.LoginRequest) (string, error)
}
