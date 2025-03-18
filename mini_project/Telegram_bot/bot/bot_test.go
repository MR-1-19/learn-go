package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleMessage(t *testing.T) {
	b := NewBotHandler()

	// Тестируем команду /start
	response := b.HandleMessage("/start", "Рауф")
	assert.Equal(t, "Привет, Рауф абый!", response)

	// Тестируем неизвестную команду
	response = b.HandleMessage("/unknown", "Рауф")
	assert.Equal(t, "", response)
}
