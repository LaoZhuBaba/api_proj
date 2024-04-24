package logger

import (
	"fmt"
	"log/slog"
)

type Logger struct {
	logger *slog.Logger
}

func (l Logger) Debug(s string, other ...any) {
	l.logger.Debug(fmt.Sprintf(s, other...))
}

func (l Logger) Error(s string, other ...any) {
	l.logger.Error(fmt.Sprintf(s, other...))
}

func (l Logger) Info(s string, other ...any) {
	l.logger.Info(fmt.Sprintf(s, other...))
}

func (l Logger) Warn(s string, other ...any) {
	l.logger.Warn(fmt.Sprintf(s, other...))
}

func New() Logger {
	// customise to suit
	return Logger{slog.Default()}
}
