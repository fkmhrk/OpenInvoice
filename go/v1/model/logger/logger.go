package logger

// Logger provides logging API
type Logger interface {
	Errorf(msg string, args ...interface{})
}
