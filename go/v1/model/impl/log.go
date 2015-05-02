package impl

import (
	"log"
	"os"
)

type logger struct {
	logger *log.Logger
}

func NewLogger() *logger {
	return &logger{
		logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

func (o *logger) Errorf(msg string, args ...interface{}) {
	o.logger.Printf(msg, args...)
}
