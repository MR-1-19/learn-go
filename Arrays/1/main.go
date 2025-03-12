package main

/* - Функция для массива
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
*/

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) (sums []int) {
	/*
		length := len(numbers)
		sums = make([]int, length)

		for i, number := range numbers {
			sums[i] = Sum(number)
		}
	*/

	for _, number := range numbers {
		sums = append(sums, Sum(number))
	}

	return sums
}

func SumAllTails(numbers ...[]int) (sums []int) {
	for _, number := range numbers {
		if len(number) > 0 {
			tail := number[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}

	}

	return sums
}

func main() {

}
