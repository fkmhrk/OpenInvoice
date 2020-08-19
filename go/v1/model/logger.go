package model

type Logger interface {
	Errorf(msg string, args ...interface{})
}
