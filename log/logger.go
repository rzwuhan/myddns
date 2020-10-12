package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger

func init() {
	//logger, _ := zap.NewProduction(zap.AddCallerSkip(1))
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := cfg.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zap.DPanicLevel))
	sugar = logger.Sugar()
}

func D(args ...interface{}) {
	sugar.Debug(args)
}

func I(args ...interface{}) {
	sugar.Info(args)
}

func W(args ...interface{}) {
	sugar.Warn(args)
}

func E(args ...interface{}) {
	sugar.Error(args)
}

func F(args ...interface{}) {
	sugar.Fatal(args)
}

func Assert(cond bool, args ...interface{}) {
	if cond {
		return
	}
	sugar.Fatal(args)
}
