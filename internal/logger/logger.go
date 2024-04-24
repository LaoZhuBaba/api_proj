package logger

import "log/slog"

type Logger struct {
	logger *slog.Logger
}

func (l Logger) Debug(s string, other ...any) {
	l.logger.Debug(s, other...)
}

func (l Logger) Error(s string, other ...any) {
	l.logger.Error(s, other...)
}

func (l Logger) Info(s string, other ...any) {
	l.logger.Info(s, other...)
}

func (l Logger) Warn(s string, other ...any) {
	l.logger.Warn(s, other...)
}

func New() Logger {
	// customise to suit
	return Logger{slog.Default()}
}
