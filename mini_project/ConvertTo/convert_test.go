package main

import (
	"testing"
)

func TestConvert(t *testing.T) {
	// Вспомогательная функция для проверки результата
	check_test := func(t *testing.T, got float64, err error, want float64) {
		t.Helper()
		if err != nil {
			t.Error(err)
		}
		epsilon := 0.01
		if (got-want) > epsilon || (want-got) > epsilon {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}

	// Тестовые случаи
	tests := []struct {
		run    string
		shkala string
		have   float64
		want   float64
	}{
		{"Проверка конвертации в цельсия", "цельсия", 77.0, 25.0},
		{"Проверка конвертации в фаренгейты", "фаренгейты", 25.0, 77.0},
		{"Проверка некорректной температуры (Цельсий)", "цельсия", -500.0, 0.0},
		{"Проверка некорректной температуры (Фаренгейт)", "фаренгейты", -300.0, 0.0},
	}

	// Запуск тестов
	for _, test := range tests {
		if test.run == "Проверка конвертации в цельсия" {
			t.Run(test.run, func(t *testing.T) {
				got, err := ConvertTo(test.have, test.shkala)
				want := test.want
				check_test(t, got, err, want)
			})
		}

		if test.run == "Проверка конвертации в фаренгейты" {
			t.Run(test.run, func(t *testing.T) {
				got, err := ConvertTo(test.have, test.shkala)
				want := test.want
				check_test(t, got, err, want)
			})
		}
	}
}
