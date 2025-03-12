package main

func RemoveDuplicates(a []int) []int {
	o := make(map[int]bool)
	b := []int{}
	for _, number1 := range a {
		if o[number1] == false {
			o[number1] = true
			b = append(b, number1)
		}
	}

	return b
}

func main() {
	a := []int{4, 2, 3, 2, 4, 5, 1}

	RemoveDuplicates(a)
}
