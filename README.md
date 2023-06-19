## Go Buzzers 📢

> Can be used to send notifications (buzzers) to your mobile or desktop devices from within go.

### Installation

```bash
go get -u github.com/mattiasbonte/gobuzzer
```

### .env

Populate the `.env` file based on the functionality you wish you use.

```bash
TELEGRAM_AUTH_TOKEN=YOURTELEGRAMAUTHTOKEN
TELEGRAM_CHAT_ID=YOURTELEGRAMCHATID
```

### Examples

```go
if err := gobuzzer.TelegramNotification(telegramAuthToken, telegramChatID, notificationMessage); err != nil {
    log.Fatalf("error sending telegram message: %v", err)
}
```

```go
if err := gobuzzer.SystemBuzz(notificationMessage, notificationType); err != nil {
    log.Fatalf("error sending system notification: %v", err)
}
```
