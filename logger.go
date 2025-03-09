package logger

import (
	"fmt"
	"github.com/Mertozturkk/NotifyLog/config"
)


// Logger is a struct that holds the notifiers for different log levels.
type Logger struct {
    Notifiers map[string][]config.INotifier
}
// NewLogger creates a new Logger instance.
func NewLogger() *Logger {
    return &Logger{
        Notifiers: make(map[string][]config.INotifier),
    }
}
// AddNotifier adds a notifier for a specific log level.
func (l *Logger) AddNotifier(level string, n config.INotifier) {
    l.Notifiers[level] = append(l.Notifiers[level], n)
}
// log logs a message and notifies the notifiers for the given log level.
func (l *Logger) log(level string, message string) {
    fmt.Printf("[%s] %s\n", level, message)
    for _, n := range l.Notifiers[level] {
        n.Notify(message)
    }
}
// INFO logs an info message.
func (l *Logger) INFO(message string) {
    l.log("INFO", message)
}
// ERROR logs an error message.
func (l *Logger) ERROR(message string) {
    l.log("ERROR", message)
}
// WARN logs a warning message.
func (l *Logger) WARN(message string) {
    l.log("WARN", message)
}
// DEBUG logs a debug message.
func (l *Logger) DEBUG(message string) {
    l.log("DEBUG", message)
}
// FATAL logs a fatal message.
func (l *Logger) FATAL(message string) {
    l.log("FATAL", message)
}
