package Telegram_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7222013854:AAEAhG41TCOMnlZyGjs6DufW1V5nnlCWZU8")
	if err != nil {
		log.Fatal("Ошибка при создании бота:", err)
	}

	bot.Debug = true // Включаем логирование запросов для отладки
	log.Println("Бот запущен:", bot.Self.UserName)

	handler := bot.NewBotHandler() // Создаем обработчик

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			response := handler.HandleMessage(update.Message.Text, update.Message.From.UserName)
			if response != "" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
				bot.Send(msg)
			}

		}
	}
}
