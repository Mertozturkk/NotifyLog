# NotifyLog

NotifyLog is a logging library that allows you to send log notifications to various platforms such as Slack, Email, and Discord. It provides a flexible and extensible way to integrate different notification channels into your logging system, ensuring that important log messages are delivered to the right people in real-time.

## Features

- Send log notifications to Slack
- Send log notifications via Email
- Send log notifications to Discord
- Easily extendable to support other notification channels
- Configurable log levels

## Installation

To install NotifyLog, use `go get`:

```sh
go get github.com/Mertozturkk/NotifyLog
```

## Usage

### Slack Notifier

```go
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
}
```

--------


### Email Notifier

```go
package main

import (
    logger "github.com/Mertozturkk/NotifyLog"
    "github.com/Mertozturkk/NotifyLog/notifier"
)

func main() {
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
```

--------

### Discord Notifier

```go
package main

import (
    logger "github.com/Mertozturkk/NotifyLog"
    "github.com/Mertozturkk/NotifyLog/notifier"
)

func main() {
    discordLogger := logger.NewLogger()
    discordNotifier := notifier.NewDiscordNotifier("https://discord.com/api/webhooks/WEBHOOK_ID/WEBHOOK_TOKEN")

    discordLogger.AddNotifier("ERROR", discordNotifier)

    discordLogger.INFO("Hello, Discord!")
    discordLogger.ERROR("Hello, Discord! ERROR")
}
```