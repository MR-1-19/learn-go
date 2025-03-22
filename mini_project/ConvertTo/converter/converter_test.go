package converter

import (
	"log"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	// Вспомогательная функция для проверки результата
	checkTest := func(t *testing.T, got float64, err error, want float64) {
		t.Helper()

		// Если ожидаемая температура ниже абсолютного нуля, то должна быть ошибка
		if want == 0.0 && err != nil {
			return // Ошибка ожидаема, тест проходит
		}

		// Если ошибка есть, но мы её не ожидали — это проблема
		if err != nil {
			t.Errorf("Не ожидалась ошибка, но получили: %v", err)
		}

		// Проверяем корректность вычисления
		epsilon := 0.01
		if (got-want) > epsilon || (want-got) > epsilon {
			t.Errorf("Получили %.2f, ожидали %.2f", got, want)
		}
	}

	// Тестовые случаи
	tests := []struct {
		name  string
		scale string
		input float64
		want  float64
	}{
		{"Конвертация в цельсия", "цельсия", 77.0, 25.0},
		{"Конвертация в фаренгейты", "фаренгейты", 25.0, 77.0},
		{"Ошибка: слишком низкая температура (Цельсий)", "цельсия", -500.0, 0.0},
		{"Ошибка: слишком низкая температура (Фаренгейт)", "фаренгейты", -300.0, 0.0},
	}

	// Запуск тестов
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertTo(tt.input, tt.scale)
			checkTest(t, got, err, tt.want)
		})
	}
}

func init() {
	file, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}
