package log

import "log"

type Logger interface {
	LogInfo(message string, fields ...string)
	LogDebug(message string, fields ...string)
	LogWarn(message string, fields ...string)
	LogError(message string, fields ...string)
}

func System(message string) {
	log.Println("[SYSTEM]", message)
}
