package logger

import (
	"log"
	"sync"
)

var (
	defaultLogger *logger
	once          sync.Once
)

type logger struct{}

func (l *logger) Log(msg string) {
	log.Println(msg)
}

func GetLogger() *logger {
	once.Do(func() {
		log.Println("initializing logger")
		defaultLogger = &logger{}
	})

	return defaultLogger
}
