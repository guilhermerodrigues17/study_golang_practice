package main

//Codewars -> https://www.codewars.com/kata/5715eaedb436cf5606000381

func SumOfPositiveRun() int {
	return sumOfPositive([]int {1, -4, 7, 12})
}

func sumOfPositive(numbers []int) int {
	var sum int

	for _, value := range numbers {
		if value > 0 {
			sum += value
		}
	}
	
	return sum
}