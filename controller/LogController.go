package controller

import (
	"github.com/gin-gonic/gin"
	"mc3_monitor_service/core"
	"mc3_monitor_service/logic"
)

//收集前端报错日志
func InsertLog(r *gin.Context) {
	source := r.DefaultQuery("source", "")
	times := r.DefaultQuery("time", "")
	token := r.DefaultQuery("token", "")
	message := r.DefaultQuery("message", "")
	if source == "" || times == "" || token == "" || message == "" {
		core.Response(r, -1, "参数不能为空", "")
		return
	}
	res, err := logic.ParseToken(source, token, message, times)
	if err != nil {
		core.Response(r, -1, err.Error(), "")
		return
	}
	if res == false {
		core.Response(r, -1, "token解析不正确", "")
		return
	}
	logic.InsertLog(source, message)
	core.Response(r, 0, "成功", "")
	return
}
