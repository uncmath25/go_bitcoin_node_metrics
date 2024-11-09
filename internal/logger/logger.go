package logger

import "log"

func BuildLogger() *log.Logger {
	return log.Default()
}
