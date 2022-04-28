package logger

type LoggerConfig struct {
	CallerKey  string // "file_name"
	TimeKey    string // "timestamp"
	CallerSkip int
}
