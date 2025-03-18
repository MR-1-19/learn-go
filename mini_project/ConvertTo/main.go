package main

import "errors"

func ConvertToCelsius(fahrenheit float64) (float64, error) {
	if fahrenheit < -459.67 {
		return 0, errors.New("температура ниже абсолютного нуля")
	}
	return (fahrenheit - 32) * 5 / 9, nil
}

func ConvertToFahrenheit(celsius float64) (float64, error) {
	if celsius < -273.15 {
		return 0, errors.New("температура ниже абсолютного нуля")
	}
	return (celsius * 9 / 5) + 32, nil
}
