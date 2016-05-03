package milog

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	debugLogger		*log.Logger
	infoLogger		*log.Logger
	errorLogger		*log.Logger
}

func New() *Logger {
	return NewWithWriters(os.Stdout, os.Stdout, os.Stderr)
}

func NewWithWriter(w io.Writer) *Logger {
	return NewWithWriters(w, w, w)
}

func NewWithFile(fp string) (*Logger, error) {
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
	    return nil, err
	}
	return NewWithWriters(f, f, f), nil
}

func NewWithWriters(debugWriter, infoWriter, errorWriter io.Writer) *Logger {
	return &Logger{
		debugLogger:	log.New(debugWriter, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLogger:		log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:	log.New(errorWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.debugLogger.Printf(format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.infoLogger.Printf(format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	color.Set(color.FgRed)
	l.errorLogger.Printf(format, args...)
	color.Unset()
}