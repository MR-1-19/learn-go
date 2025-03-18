package bot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleMessage(t *testing.T) {
	b := NewBotHandler()

	// Тестируем команду /start
	response := b.HandleMessage("/start", "Рауф")
	assert.Equal(t, "Привет, Рауф абый!", response)

	// Тестируем неизвестную команду
	response = b.HandleMessage("/unknown", "Рауф")
	assert.Equal(t, "", response)

	// Тестируем команду /help
	response = b.HandleMessage("/help", "Рауф")
	assert.Equal(t, "Я помогаю учителям! 📚\n\nКоманды:\n/start - Приветствие\n/help - Список команд\n/schedule - Посмотреть расписание", response)

	// Тестируем команду /schedule
	expectedSchedule := "📅 Расписание занятий на неделю:\n\n🟢 Понедельник: Основы робототехники\n🔵 Вторник: Программирование\n🟠 Среда: Механика\n🟣 Четверг: Практика\n🔴 Пятница: Подготовка к соревнованиям"
	response = b.HandleMessage("/schedule", "Рауф")
	assert.Equal(t, expectedSchedule, response)
}
