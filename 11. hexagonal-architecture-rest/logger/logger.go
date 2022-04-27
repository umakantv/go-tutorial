package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Init(loggerConfig LoggerConfig) {
	var err error

	eConfig := zap.NewProductionEncoderConfig()
	eConfig.CallerKey = loggerConfig.CallerKey
	eConfig.TimeKey = loggerConfig.TimeKey
	eConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.NewProductionConfig()
	config.EncoderConfig = eConfig

	// log, err = zap.NewProduction(zap.AddCallerSkip(1))
	log, err = config.Build(zap.AddCallerSkip(loggerConfig.CallerSkip))

	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
