package logger

import (
	"fmt"
	"github.com/Mertozturkk/NotifyLog/config"
)


type Logger struct {
    Notifiers map[string][]config.INotifier
}

func NewLogger() *Logger {
    return &Logger{
        Notifiers: make(map[string][]config.INotifier),
    }
}

func (l *Logger) AddNotifier(level string, n config.INotifier) {
    l.Notifiers[level] = append(l.Notifiers[level], n)
}

func (l *Logger) log(level string, message string) {
    fmt.Printf("[%s] %s\n", level, message)
    for _, n := range l.Notifiers[level] {
        n.Notify(message)
    }
}
func (l *Logger) INFO(message string) {
    l.log("INFO", message)
}

func (l *Logger) ERROR(message string) {
    l.log("ERROR", message)
}

func (l *Logger) WARN(message string) {
    l.log("WARN", message)
}

func (l *Logger) DEBUG(message string) {
    l.log("DEBUG", message)
}

func (l *Logger) FATAL(message string) {
    l.log("FATAL", message)
}
