package core

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, code int, message interface{}, data interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  message,
		"data": data,
	})
}
