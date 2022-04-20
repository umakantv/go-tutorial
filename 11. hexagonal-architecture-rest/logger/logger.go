package logger

import "go.uber.org/zap"

var log *zap.Logger

func init() {
	var err error
	log, err = zap.NewProduction(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
