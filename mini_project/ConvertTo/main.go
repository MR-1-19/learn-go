package main

import (
	"ConvertTo/converter"
	"fmt"
	"log"
	"os"
)

func init() {
	// Открываем файл логов
	file, err := os.OpenFile("conversions.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Ошибка открытия файла логов:", err)
	}
	log.SetOutput(file) // Записываем логи в файл
}

func main() {
	var temp float64
	var scale string

	// Ввод температуры
	fmt.Print("Введите температуру: ")
	_, err := fmt.Scanln(&temp)
	if err != nil {
		log.Println("Ошибка ввода температуры:", err)
		fmt.Println("Ошибка ввода температуры. Попробуйте снова.")
		return
	}

	// Ввод шкалы (Цельсия или Фаренгейты)
	fmt.Print("Введите шкалу (цельсия/фаренгейты): ")
	_, err = fmt.Scanln(&scale)
	if err != nil {
		log.Println("Ошибка ввода шкалы:", err)
		fmt.Println("Ошибка ввода шкалы. Попробуйте снова.")
		return
	}

	// Конвертация температуры
	result, err := converter.ConvertTo(temp, scale)
	if err != nil {
		log.Println("Ошибка преобразования:", err)
		fmt.Println("Ошибка:", err)
		return
	}

	// Вывод результата
	log.Printf("Конвертация: %.2f %s -> %.2f\n", temp, scale, result)
	fmt.Printf("Результат: %.2f\n", result)
}
