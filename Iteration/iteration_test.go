package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 5)
	expected := "aaaaa"

	if repeat != expected {
		t.Errorf("Получили %s, хотели %s", repeat, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	s := Repeat("a", 5)
	fmt.Println(s)
	// Output: aaaaa
}
