package logic

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mc3_monitor_service/core"
	"mc3_monitor_service/helper"
)

//解析token是否正确
func ParseToken(source, token, message, times string) (bool, error) {
	secret := fmt.Sprintf("secret.%s", source)
	secret = viper.GetString(secret)
	if secret == "" {
		return false, errors.New("密钥不存在")
	}
	key := fmt.Sprintf("%s+%s+%s", secret, message, times)
	newToken := helper.Md5(key)
	if newToken != token {
		core.Logger.Info("token_error", zap.Any("key", key))
		return false, nil
	}
	return true, nil
}

//收集前端报错日志
func InsertLog(source, message string) {
	source = "front_error"
	core.Logger.Info(source, zap.Any("data", message))
}
