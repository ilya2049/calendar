package log

import "log"

func System(message string) {
	log.Println("[SYSTEM]", message)
}
