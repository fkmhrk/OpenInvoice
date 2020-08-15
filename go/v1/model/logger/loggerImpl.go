package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type logger struct {
	logger *log.Logger
}

// New creates instance
func New() Logger {
	return &logger{
		logger: log.New(os.Stdout, "", 0),
	}
}

func (o *logger) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = strings.Replace(msg, ":", "|", -1)
	msg = fmt.Sprintf("msg:%s\tdate:%s", msg, o.createDate())
	o.logger.Print(msg)
}

func (o *logger) createDate() string {
	t := time.Now()
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	return fmt.Sprintf("%04d/%02d/%02d %02d-%02d-%02d", year, month, day, hour, min, sec)
}
