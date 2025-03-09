package notifier

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// SlackNotifier is a struct that holds the configuration for sending messages to Slack.
type SlackNotifier struct {
	WebHookUrl string
	UserName   string
	Channel    string
	IconEmoji  string
	TimeOut    time.Duration
}

// Notify sends a message to the configured Slack webhook.
func (n SlackNotifier) Notify(message string) {
	slackMessage := SlackMessage{
		Username:  n.UserName,
		Channel:   n.Channel,
		Text:      message,
		IconEmoji: n.IconEmoji,
	}
	sendHttpRequest(n.WebHookUrl, slackMessage, n.TimeOut)
}

// SlackMessage is a struct that represents a message to be sent to Slack.
type SlackMessage struct {
	Username    string       `json:"username,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Attachment is a struct that represents an attachment in a Slack message.
type Attachment struct {
	Fallback   string `json:"fallback,omitempty"`
	Color      string `json:"color,omitempty"`
	Pretext    string `json:"pretext,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
}

// sendHttpRequest sends an HTTP request to the given URL with the given Slack message.
func sendHttpRequest(url string, message SlackMessage, timeout time.Duration) error {
	slackBody, _ := json.Marshal(message)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	if buf.String() != "ok" {
		return errors.New("non-ok response returned from Slack")
	}
	return nil
}
