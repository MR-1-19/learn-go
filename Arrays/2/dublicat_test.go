package main

import (
	"reflect"
	"testing"
)

func TestDublicat(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		want []int
	}{
		{"no duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"with duplicates", []int{1, 2, 2, 3, 4, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"empty slice", []int{}, []int{}},
		{"all duplicates", []int{1, 1, 1, 1}, []int{1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := RemoveDuplicates(test.a)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Получили %v; хотели %v", got, test.want)
			}
		})
	}
}
