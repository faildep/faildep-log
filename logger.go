package log

import slog "log"

type Logger interface {
	Error(args ...interface{})
	Warning(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

var _ Logger = &StdLogger{}

type StdLogger struct{}

func (l *StdLogger) Error(args ...interface{}) {
	args = append([]interface{}{"[Error]"}, args...)
	slog.Println(args...)
}

func (l *StdLogger) Warning(args ...interface{}) {
	args = append([]interface{}{"[Warning]"}, args...)
	slog.Println(args...)
}

func (l *StdLogger) Info(args ...interface{}) {
	args = append([]interface{}{"[Info]"}, args...)
	slog.Println(args...)
}

func (l *StdLogger) Debug(args ...interface{}) {
	args = append([]interface{}{"[Debug]"}, args...)
	slog.Println(args...)
}
