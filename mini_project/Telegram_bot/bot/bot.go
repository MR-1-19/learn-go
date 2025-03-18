package bot

import "fmt"

type BotHandler struct{}

func NewBotHandler() *BotHandler {
	return &BotHandler{}
}

func (b *BotHandler) HandleMessage(msg string, userName string) string {
	if msg == "/start" {
		return fmt.Sprintf("Привет, %s абый!", userName)
	}
	return ""
}
