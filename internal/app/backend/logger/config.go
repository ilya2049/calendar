package logger

type LogLevel = int8

const (
	DebugLogLevel LogLevel = -1
	InfoLogLevel  LogLevel = 0
	WarnLogLevel  LogLevel = 1
	ErrorLogLevel LogLevel = 2
)

type Config struct {
	LogLevel LogLevel
}
