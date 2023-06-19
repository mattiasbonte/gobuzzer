## Go Buzzers 📢

> Can be used to send notifications (buzzers) to your mobile or desktop devices from within go.

### Installation

```bash
go get -u github.com/mattiasbonte/gobuzzer
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
