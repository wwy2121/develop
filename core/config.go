package core

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
)

var (
	Logger         *zap.Logger
)

//加载配置文件到vipper中
func LoadConfigFile(cfgFile string) error {
	viper.SetConfigFile(cfgFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//初始化配置内容
func InitConfile() (err error) {
	//加载zap日志
	Logger, err = NewZap()
	return nil
}
