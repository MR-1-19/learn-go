package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Создаем новое приложение
	a := app.New()

	// Создаем новое окно с заголовком "Input Example"
	w := a.NewWindow("Input Example")

	// Создаем поле для ввода текста
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter text...")

	// Создаем текстовую метку
	output := widget.NewLabel("")

	// Создаем кнопку
	button := widget.NewButton("Submit", func() {
		output.SetText("You entered: " + entry.Text)
	})

	// Размещаем элементы в окне
	w.SetContent(container.NewVBox(
		entry,
		button,
		output,
	))

	// Показываем окно и запускаем приложение
	w.ShowAndRun()
}
