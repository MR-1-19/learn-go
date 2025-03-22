package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"learn-go/mini_project/ConvertTo/converter"
	"log"
	"os"
	"strconv"
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
	// Создаём новое приложение
	a := app.New()

	// Создаём окно
	w := a.NewWindow("Конвертер температуры")

	inputTemp := widget.NewEntry()
	inputTemp.SetPlaceHolder("Введите температуру")

	scale := widget.NewSelect([]string{"цельсия", "фаренгейты"}, nil)
	scale.SetSelected("цельсия") // Значение по умолчанию

	label := widget.NewLabel("Результат")

	var convertButton *widget.Button

	convertButton = widget.NewButton("Конвертировать", func() {
		// Получаем данные от пользователя
		tempStr := inputTemp.Text
		scaleSelect := scale.Selected

		// Преобразуем ввод в число
		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			log.Fatal(err)
			label.SetText("Ошибка: Введите корректное число")
			return
		}

		// Выполняем конвертацию
		converted, err := converter.ConvertTo(temp, scaleSelect)
		if err != nil {
			log.Fatal(err)
			label.SetText(fmt.Sprintf("Ошибка: %s", err.Error()))
			return
		}

		label.SetText(fmt.Sprintf("Результат: %.2f", converted))

		log.Printf("Конвертация: %.2f %s -> %.2f\n", temp, scaleSelect, converted)

	})

	// Добавляем кнопку в окно
	w.SetContent(container.NewVBox(
		widget.NewLabel("Конвертер температуры"),
		inputTemp,
		scale,
		convertButton,
		label,
	))
	// Устанавливаем размер окна
	w.Resize(fyne.NewSize(300, 200))

	// Запускаем приложение
	w.ShowAndRun()
}
