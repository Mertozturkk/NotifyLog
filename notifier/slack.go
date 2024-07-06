package notifier

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type SlackNotifier struct {
    WebHookUrl string
    UserName   string
    Channel    string
    TimeOut    time.Duration
}

func (n SlackNotifier) Notify(message string) {
    slackMessage := SlackMessage{
        Username: n.UserName,
        Channel:  n.Channel,
        Text:     message,
    }
    sendHttpRequest(n.WebHookUrl, slackMessage, n.TimeOut)
}

type SlackMessage struct {
    Username    string       `json:"username,omitempty"`
    IconEmoji   string       `json:"icon_emoji,omitempty"`
    Channel     string       `json:"channel,omitempty"`
    Text        string       `json:"text,omitempty"`
    Attachments []Attachment `json:"attachments,omitempty"`
}

type Attachment struct {
    Fallback   string  `json:"fallback,omitempty"`
    Color      string  `json:"color,omitempty"`
    Pretext    string  `json:"pretext,omitempty"`
    AuthorName string  `json:"author_name,omitempty"`
}

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