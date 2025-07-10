package main

//Codewars -> https://www.codewars.com/kata/53da3dbb4a5168369a0000fe

func OddOrEvenRun() string {
	return oddOrEven(2)
}

func oddOrEven(n int) string {
	if n % 2 == 0 {return "Even"} 
	return "Odd"
}
