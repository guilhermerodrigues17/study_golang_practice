package main

//Codewars -> https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad

func InvertValuesRun() []int {
	return invertValues([]int{-1, 2, -3, 4, -5})
}

func invertValues(arr []int) (result []int) {
	for _, n := range arr {
		result = append(result, -n)
	}
	return
}