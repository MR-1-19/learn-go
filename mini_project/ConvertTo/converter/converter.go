package converter

import (
	"errors"
	"strings"
)

const (
	minFar = -459.67 // Абсолютный ноль в Фаренгейтах
	minCel = -273.15 // Абсолютный ноль в Цельсиях
)

// ConvertTo выполняет конвертацию температуры в заданную шкалу.
func ConvertTo(temp float64, scale string) (float64, error) {
	scale = strings.TrimSpace(strings.ToLower(scale)) // Убираем пробелы и приводим к нижнему регистру
	switch scale {
	case "цельсия":
		return ConvertToCelsius(temp)
	case "фаренгейты":
		return ConvertToFahrenheit(temp)
	default:
		return 0, errors.New("неизвестная шкала температуры")
	}
}

// ConvertToCelsius переводит Фаренгейты в Цельсии.
func ConvertToCelsius(fahrenheit float64) (float64, error) {
	if fahrenheit < minFar {
		return 0.0, errors.New("температура ниже абсолютного нуля (Фаренгейты)")
	}
	return (fahrenheit - 32) * 5 / 9, nil
}

// ConvertToFahrenheit переводит Цельсии в Фаренгейты.
func ConvertToFahrenheit(celsius float64) (float64, error) {
	if celsius < minCel {
		return 0.0, errors.New("температура ниже абсолютного нуля (Цельсии)")
	}
	return (celsius*9/5 + 32), nil
}
