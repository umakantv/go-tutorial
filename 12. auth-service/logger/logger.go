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

type Level = zapcore.Level

const (
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel Level = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel // 5
	DebugLevel Level = zap.DebugLevel // -1
)

var (
	Skip       = zap.Skip
	Binary     = zap.Binary
	Bool       = zap.Bool
	Boolp      = zap.Boolp
	ByteString = zap.ByteString
	Float64    = zap.Float64
	Float64p   = zap.Float64p
	Float32    = zap.Float32
	Float32p   = zap.Float32p
	Durationp  = zap.Durationp
	Any        = zap.Any

	// Info   = std.Info
	// Warn   = std.Warn
	// Error  = std.Error
	// DPanic = std.DPanic
	// Panic  = std.Panic
	// Fatal  = std.Fatal
	// Debug  = std.Debug
)
