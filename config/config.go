package config

import (
	"github.com/Mertozturkk/NotifyLog/notifier"

)
type INotifier interface {
	Notify(message string)
}


type Config struct {
	SlackNotifier notifier.SlackNotifier
}