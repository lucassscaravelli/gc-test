package logservice

import (
	"fmt"
)

type LogSevice struct {
	packageName string
}

func NewLogService(packageName string) *LogSevice {
	return &LogSevice{packageName}
}

func (l *LogSevice) Error(msg interface{}) {
	fmt.Println("["+l.packageName+"] ", " - ERROR -", msg)
}

func (l *LogSevice) Info(msg interface{}) {
	fmt.Println("["+l.packageName+"] ", " - INFO -", msg)
}
