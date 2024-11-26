package controller

import "github.com/gin-gonic/gin"



func LoadIndex(e *gin.Engine) {
	e.GET("/index.html", welcome)
	e.GET("/", welcome)
}

func welcome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
