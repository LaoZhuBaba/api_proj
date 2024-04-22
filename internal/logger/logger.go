package logger

import "log"

func LogOutput(s string) {
	log.Print(s)
}

type Logger interface {
	Log(string)
}
