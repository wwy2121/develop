package controller

import (
	"github.com/gin-gonic/gin"
)

//路由
func InitRoute(r *gin.Engine) {
	r.GET("/log/insert", InsertLog)
}
