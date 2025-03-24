package notifier

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)


type WebhookNotifier struct {
	URL string
	Method string
	Headers map[string]string
	TimeOut time.Duration
	Retries int
}

func NewWebhookNotifier(url string, method string, headers map[string]string, timeout time.Duration, retries int) *WebhookNotifier {
	return &WebhookNotifier{
		URL: url,
		Method: method,
		Headers: headers,
		TimeOut: timeout,
		Retries: retries,
	}
}

func (w *WebhookNotifier) Notify(message string) error{
	payload := map[string]string{"message": message}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for i := 0; i < w.Retries; i ++{
		err = w.sendRequest(jsonPayload)
		if err == nil {
			return nil
		}
		time.Sleep(2 * time.Second)

	}

	return err
}

func (w *WebhookNotifier) sendRequest(jsonPayload []byte) error {
	req,err := http.NewRequest(w.Method, w.URL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range w.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: w.TimeOut,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode <200 || resp.StatusCode >= 300 {
		return err
	}

	return nil
}