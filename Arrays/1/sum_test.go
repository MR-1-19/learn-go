package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	/*
		t.Run("Проверка массива из 5 элементов", func(t *testing.T) {
			numbers := [5]int{1, 2, 3, 4, 5}

			got := Sum(numbers)
			want := 15

			if got != want {
				t.Errorf("Получили %d хотели %d given, %v", got, want, numbers)
			}
		})
	*/
	t.Run("Проверка срезов", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Получили %d хотели %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Получили: %v, хотели: %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Получили: %v, хотели: %v", got, want)
		}
	}

	t.Run("Сумма хвостов среза", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("Передача пустого среза", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
