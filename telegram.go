package gobuzzer

import (
	"net/http"
	"net/url"
)

type TelegramMsg struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

// TelegramNotification sends a message to a Telegram chat.
func TelegramNotification(authToken string, chatID string, message string) error {
	apiURL := "https://api.telegram.org/bot" + authToken + "/sendMessage"

	values := url.Values{}
	values.Add("chat_id", chatID)
	values.Add("text", message)
	values.Add("parse_mode", "MarkdownV2")

	resp, err := http.Get(apiURL + "?" + values.Encode())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
