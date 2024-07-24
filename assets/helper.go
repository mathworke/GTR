package assets

import (
	"log"
	"os"
)

const (
	LOG_PATH = "logs/.log"
)

type Logger struct {
	Log *os.File

	logging bool

	logInfo  *log.Logger
	logError *log.Logger
}

func NewLogger(logging bool) *Logger {
	os.Mkdir("logs", 066)
	file, err := os.OpenFile(LOG_PATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return &Logger{
		Log:      file,
		logging:  logging,
		logInfo:  log.New(file, "[INFO]\t", log.Ldate|log.Ltime),
		logError: log.New(file, "[ERROR]\t", log.Ldate|log.Ltime),
	}
}

func (l *Logger) LogIngo(format string, args ...any) {
	if l.logging {
		l.logInfo.Printf(format, args...)
	}
}

func (l *Logger) LogError(format string, args ...any) {
	if l.logging {
		l.logError.Printf(format, args...)
	}
}

func (l *Logger) PANIC(format string, args ...any) {
	if l.logging {
		l.logError.Panicf(format, args...)
	}
}
