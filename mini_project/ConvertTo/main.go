package main

import (
	"errors"
	"fmt"
)

const minfar = -459.67
const mincel = -273.15

func ConvertTo(temp float64, scale string) (float64, error) {
	switch scale {
	case "цельсия":
		return ConvertToCelsius(temp)
	case "фаренгейты":
		return ConvertToFahrenheit(temp)
	default:
		return 0, errors.New("Ошибка в шкале")
	}
}

func ConvertToCelsius(fahrenheit float64) (float64, error) {
	if fahrenheit < minfar {
		return 0, errors.New("температура ниже абсолютного нуля")
	}
	return (fahrenheit - 32) * 5 / 9, nil
}

func ConvertToFahrenheit(celsius float64) (float64, error) {
	if celsius < mincel {
		return 0, errors.New("температура ниже абсолютного нуля")
	}
	return (celsius * 9 / 5) + 32, nil
}

func main() {
	var temp float64
	var scale string

	// Ввод температуры
	fmt.Print("Введите температуру: ")
	_, err := fmt.Scanf("%f", &temp)
	if err != nil {
		fmt.Println("Ошибка ввода температуры:", err)
		return
	}

	// Ввод шкалы
	fmt.Print("Введите шкалу (цельсия/фаренгейты): ")
	_, err = fmt.Scanf("%s", &scale)
	if err != nil {
		fmt.Println("Ошибка ввода шкалы:", err)
		return
	}

	ConvertTo(142, fmt.Scanf(""))
}
