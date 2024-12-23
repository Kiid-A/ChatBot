package controller

import (
	"ChatBot/internal/model"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 负责API的路由，定向到指定服务方法

var nowUser *model.User

func GetNowUser() *model.User {
	return nowUser
}

type UserController struct {
	DB *gorm.DB
}

func (uc *UserController) LoadUser(e *gin.Engine) {
	e.GET("/login.html", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", gin.H{})
	})
	e.GET("/register.html", func(ctx *gin.Context) {
		ctx.HTML(200, "register.html", gin.H{})
	})
	e.GET("/user_info.html", func(ctx *gin.Context) {
		ctx.HTML(200, "user_info.html", gin.H{})
	})

	e.GET("/user/:id/info", uc.UserInfo)
	e.POST("/user/login", uc.Login)
	e.POST("/user/register", uc.Register)
	e.PUT("/user/:id", uc.Modify)
}

func (uc *UserController) UserInfo(ctx *gin.Context) {
	var reqUser model.User
	id := ctx.Param("id")
	if err := uc.DB.Table("users").Where("id = ?", id).Find(&reqUser).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Invalid user id " + err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"user": reqUser,
	})
}

func (uc *UserController) Modify(ctx *gin.Context) {
	var reqUser model.User
	ctx.BindJSON(&reqUser)
	id := ctx.Param("id")
	reqUser.Id = id
	if err := uc.DB.Table("users").Where("id = ?", id).Update("name", reqUser.Name).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Invalid area name " + err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"msg": "modify successfully"})
}

func (uc *UserController) Login(ctx *gin.Context) {
	uc.DB.Table("users")
	var req model.User
	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	fmt.Println("guest ", req.Username, req.Passwd)

	var user model.User
	uc.DB.Where("username = ?", req.Username).First(&user)
	if user.ID != 0 {
		ctx.JSON(500, gin.H{"error": "username already exist"})
		return
	}

	if flag := strings.Compare(req.Passwd, user.Passwd); flag != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})
		return
	}

	nowUser = &user

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) Register(ctx *gin.Context) {
	uc.DB.Table("users")
	var reqUser model.User
	if err := ctx.BindJSON(&reqUser); err != nil {
		return
	}

	if err := uc.DB.Create(&reqUser).Error; err != nil {
		ctx.JSON(500, "cannot create user "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "register successfully",
	})

	var cnt int64
	if err := uc.DB.Model(&model.User{}).Count(&cnt).Error; err != nil {
		panic("failed to count users")
	}
	reqUser.Id = strconv.Itoa(int(cnt))
	uc.DB.Model(&reqUser).Update("id", reqUser.Id)

	nowUser = &reqUser
}
