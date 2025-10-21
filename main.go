package main

import (
	"fmt"
)

type Logger struct{}

func (l *Logger) Log(message string) {
	fmt.Println("Message:", message)
}

var instance *Logger

func GetLoggerInstance() *Logger {
	if instance == nil {
		fmt.Println("Creating Logger instance...")
		instance = &Logger{}
	}
	return instance
}

func main() {
	logger1 := GetLoggerInstance()
	logger2 := GetLoggerInstance()

	logger1.Log("Application started")
	logger2.Log("User logged in")

	fmt.Println("Are both loggers same?", logger1 == logger2)
}
