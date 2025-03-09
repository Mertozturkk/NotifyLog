package main

import (
	"time"

	logger "github.com/Mertozturkk/NotifyLog"
	"github.com/Mertozturkk/NotifyLog/notifier"
)

func main() {
	log := logger.NewLogger()

	slackNotifier := notifier.SlackNotifier{
		WebHookUrl: "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXX",
		UserName:   "NotifyLog",
		Channel:    "test",
		TimeOut:    5 * time.Second,
		IconEmoji:  ":ghost:",
	}

	log.AddNotifier("ERROR", &slackNotifier)

	log.INFO("Hello, Slack!")
	log.ERROR("Hello, Slack! ERROR")

	mailLogger := logger.NewLogger()
	emailNotifier := notifier.EmailNotifier{
		SMTPServer: "smtp.gmail.com",
		Port:       "587",
		Username:   "test@gmail.com",
		Password:   "THIS_IS_MY_PASSWORD",
		To:         []string{"test2@gmail.com"},
	}

	mailLogger.AddNotifier("ERROR", &emailNotifier)

	mailLogger.INFO("Hello, Email!")
	mailLogger.ERROR("Hello, Email! ERROR")

}
