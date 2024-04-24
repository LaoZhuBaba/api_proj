package logger

import "log"

func LogOutput(s string, other ...any) {
	log.Printf(s, other...)
}
