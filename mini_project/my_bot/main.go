package main

import (
	"log"
	"my_bot/bot"
	"my_bot/database" // Подключаем наш новый модуль

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	database.InitDB() // Инициализируем базу данных

	botAPI, err := tgbotapi.NewBotAPI("7222013854:AAEAhG41TCOMnlZyGjs6DufW1V5nnlCWZU8")
	if err != nil {
		log.Fatal("Ошибка при создании бота:", err)
	}

	botAPI.Debug = true
	log.Println("Бот запущен:", botAPI.Self.UserName)

	handler := bot.NewBotHandler()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := botAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			response := handler.HandleMessage(update.Message.Text, update.Message.Chat.FirstName)
			if response != "" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
				msg.ParseMode = "Markdown" // Включаем Markdown
				botAPI.Send(msg)
			}
		}
	}
}
