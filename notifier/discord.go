package notifier

import (
    "bytes"
    "encoding/json"
    "net/http"
    "fmt"
)

// DiscordNotifier is a struct that holds the configuration for sending messages to Discord.
type DiscordNotifier struct {
    WebhookURL string
}

// NewDiscordNotifier creates a new DiscordNotifier instance.
func NewDiscordNotifier(webhookURL string) *DiscordNotifier {
    return &DiscordNotifier{
        WebhookURL: webhookURL,
    }
}

// Notify sends a message to the configured Discord webhook.
func (d *DiscordNotifier) Notify(message string) error {
    payload := map[string]string{"content": message}
    jsonPayload, err := json.Marshal(payload)
    if err != nil {
        return err
    }

    resp, err := http.Post(d.WebhookURL, "application/json", bytes.NewBuffer(jsonPayload))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("failed to send message to Discord: %s", resp.Status)
    }

    return nil
}