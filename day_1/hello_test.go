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

		got := Hello("Chris")
		want := "Hello Chris"

		func_test(t, got, want)
	})

	t.Run("Передача пустой строчки", func(t *testing.T) {

		got := Hello("")
		want := "Hello world"

		func_test(t, got, want)
	})

}
