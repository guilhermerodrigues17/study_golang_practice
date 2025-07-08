package main

// Codewars -> https://www.codewars.com/kata/55d24f55d7dd296eb9000030

func SummationRun() int {
	return summation(10)
}

func summation(n int) int {
	var value int

	for i := 1; i <= n; i++ {
		value += i
	}
	return value
}
