package config

import (
	"github.com/Mertozturkk/NotifyLog/notifier"

)
// INotifier is an interface that defines the Notify method.
type INotifier interface {
	Notify(message string)
}


// Config is a struct that holds the configuration for different notifiers.
type Config struct {
	SlackNotifier notifier.SlackNotifier
}