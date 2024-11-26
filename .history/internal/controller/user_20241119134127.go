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

var nowUser *model.User

func GetNowUser() *model.User {
	return nowUser
}

type UserController strut {
	DB *gorm.DB
}

func (u *UserController) LoadUser(e *gin.Engine) {
	e.GET("/login.html", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{})
	})
	e.GET("/register.html", func(c *gin.Context) {
		c.HTML(200, "register.html", gin.H{})
	})
	e.GET("/user_info.html", func(c *gin.Context) {
		c.HTML(200, "user_info.html", gin.H{})
	})

	e.GET("/api/user/:id/info", u.userInfo)
	e.POST("/api/login", u.login)
	e.POST("/api/register", u.register)
	e.PUT("/api/user/:id", u.modify)
}

func (u *UserController) userInfo(c *gin.Context) {
	var reqUser model.User
	id := c.Param("id")
	if err := u.DB.Table("users").Where("id = ?", id).Find(&reqUser).Error; err != nil {
		c.JSON(500, gin.H{"error": "Invalid user id " + err.Error()})
		return
	}

	var articles []model.Article
	if err := u.DB.Table("articles").Where("author_id = ?", id).Order("updated_at desc").Find(&articles).Error; err != nil {
		c.JSON(500, gin.H{"error": "Invalid user id " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"articles": articles,
		"user":     reqUser,
	})
}

func (u *UserController) modify(c *gin.Context) {
	var reqUser model.User
	c.BindJSON(&reqUser)
	id := c.Param("id")
	reqUser.Id = id
	if err := u.DB.Table("users").Where("id = ?", id).Update("name", reqUser.Name).Error; err != nil {
		c.JSON(500, gin.H{"error": "Invalid area name " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "modify sucessfully"})
}

func (u *UserController) login(c *gin.Context) {
	u.DB.Table("users")
	username := c.Query("username")
	passwd := c.Query("passwd")

	fmt.Println("guest ", username, passwd)

	var user model.User
	u.DB.Where("username = ?", username).First(&user)
	if user.ID != 0 {
		c.JSON(500, gin.H{"error": "username already exist"})
		return
	}

	if flag := strings.Compare(passwd, user.Passwd); flag != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})
		return
	}

	nowUser = &user

	c.JSON(http.StatusOK, user)
}

func (u *UserController) register(c *gin.Context) {
	u.DB.Table("users")
	var reqUser model.User
	c.BindJSON(&reqUser)

	if err := u.DB.Create(&reqUser).Error; err != nil {
		c.JSON(500, "cannot create user "+err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "register sucessfully",
	})

	var cnt int64
	if err := u.DB.Model(&model.User{}).Count(&cnt).Error; err != nil {
		panic("failed to count users")
	}
	reqUser.Id = strconv.Itoa(int(cnt))
	u.DB.Model(&reqUser).Update("id", reqUser.Id)

	nowUser = &reqUser
}
