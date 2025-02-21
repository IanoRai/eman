package log

import "fmt"

type LoggerLevel int

// the lower the level, the more logs.
const (
	DebugLevel LoggerLevel = iota
	InfoLevel
	ErrorLevel
)

var currentLevel = InfoLevel

func SetLevel(_level LoggerLevel) {
	currentLevel = _level
}

func Log(level LoggerLevel, args ...interface{}) {
	if currentLevel <= level {
		fmt.Println(args...)
	}
}

func Logf(level LoggerLevel, message string, args ...interface{}) {
	if currentLevel <= level {
		fmt.Printf(message, args...)
	}
}
