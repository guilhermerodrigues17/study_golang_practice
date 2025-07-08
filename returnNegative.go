package main

//Codewars -> https://www.codewars.com/kata/55685cd7ad70877c23000102

func ReturnNegativeRun() int {
	return returnNegative(46)
}

func returnNegative(x int) int {
	if x <= 0 {
		return x
	}

	return -x
}