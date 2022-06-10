package logger

import (
	"log"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Error(err error) {
	log.Println(err.Error())
}

func (l *Logger) Info(msg string) {
	log.Println(msg)
}
