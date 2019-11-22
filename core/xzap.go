package core

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
	"time"
)

const errMsg = "log file not found"

func NewZap() (log *zap.Logger, err error) {
	if !viper.IsSet("log") {
		err = errors.New(errMsg)
		return
	}
	var zcfg zap.Config
	zcfg = zap.NewProductionConfig()
	zcfg.OutputPaths = viper.GetStringSlice("log.outputPaths")
	zcfg.Level = zap.NewAtomicLevelAt(zapcore.Level(viper.GetInt32("log.level")))
	zcfg.Development = viper.GetBool("log.development")
	zcfg.EncoderConfig.EncodeTime = TimeEncoder
	zcfg.DisableStacktrace = viper.GetBool("log.disablestacktrace")
	log, err = zcfg.Build()
	return
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
