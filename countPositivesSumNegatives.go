package main

func CountPositiveSumNegativeRun() []int {
	return countPositiveSumNegative([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15})
}

func countPositiveSumNegative(numbers []int) []int {
	var count, sum int

	for _, n := range numbers {
		if n > 0 {
			count++
		}
		if n < 0 {
			sum += n
		}
	}

	return []int {count, sum}
}