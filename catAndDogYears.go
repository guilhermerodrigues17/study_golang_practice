package main

//Codewars -> https://www.codewars.com/kata/5a6663e9fd56cb5ab800008b

func CatAndDogYearsRun() [3]int {
	return catAndDogYears(10)
}

func catAndDogYears(years int) (result [3]int) {
	switch years {
		case 1: result = [3]int{1, 15, 15}
		case 2: result = [3]int{2, 24, 24}
		default: result = [3]int{years, 24 + 4 * (years - 2), 24 + 5 * (years - 2)}
	}
	return
}