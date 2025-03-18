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
	if msg == "/help" {
		return "Я помогаю учителям! 📚\n\nКоманды:\n/start - Приветствие\n/help - Список команд\n/schedule - Посмотреть расписание"
	}
	if msg == "/schedule" {
		return "📅 *Расписание занятий на неделю:*\n\n" +
			"🟢 *Понедельник:* Основы робототехники\n" +
			"🔵 *Вторник:* Программирование\n" +
			"🟠 *Среда:* Механика\n" +
			"🟣 *Четверг:* Практика\n" +
			"🔴 *Пятница:* Подготовка к соревнованиям"
	}
	return ""
}
