package gobuzzer

import (
	"net/http"
	"net/url"
	"strconv"
)

type TelegramNotificationOptions struct {
	ParseMode             string // "MarkdownV2" | "HTML" | "Markdown"
	DisableWebPagePreview bool
	DisableNotification   bool
}

// TelegramNotification sends a message to a Telegram chat.
// https://core.telegram.org/bots/api#sendmessage
func TelegramNotification(authToken string, chatId string, message string, opts *TelegramNotificationOptions) error {
	apiURL := "https://api.telegram.org/bot" + authToken + "/sendMessage"
	values := url.Values{}
	values.Add("chat_id", chatId)
	values.Add("text", message)

	// Optional Params
	if &opts.ParseMode != nil {
		values.Add("parse_mode", opts.ParseMode)
	}
	if &opts.DisableNotification != nil {
		values.Add("disable_notification", strconv.FormatBool(opts.DisableNotification))
	}
	if &opts.DisableWebPagePreview != nil {
		values.Add("disable_web_page_preview", strconv.FormatBool(opts.DisableWebPagePreview))
	}

	resp, err := http.Get(apiURL + "?" + values.Encode())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
