package singleton

import (
	"fmt"
	"sync"
)

type Logger struct{}

var instance *Logger
var once sync.Once

func GetInstance() *Logger {
	once.Do(func() {
		instance = &Logger{}
	})
	return instance
}

func (l *Logger) Info(msg string) {
	fmt.Println("[info]", msg)
}

func (l *Logger) Error(msg string) {
	fmt.Println("[error]", msg)
}
