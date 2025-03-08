package main

import "testing"

func TestHello(t *testing.T) {

	func_test := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Получили %q хотели %q", got, want)
		}
	}

	t.Run("Говорить привет для людей", func(t *testing.T) {

		got := Hello("Chris", "English")
		want := "Hello Chris"

		func_test(t, got, want)
	})

	t.Run("Передача пустой строчки", func(t *testing.T) {

		got := Hello("", "English")
		want := "Hello world"

		func_test(t, got, want)
	})

	t.Run("Проверка испонского языка", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola Chris"
		func_test(t, got, want)

	})

	t.Run("Проверка французского языка", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour Chris"
		func_test(t, got, want)

	})

	t.Run("Проверка английского языка", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello Chris"
		func_test(t, got, want)

	})

}
