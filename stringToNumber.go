package main

import "strconv"

//Codewars -> https://www.codewars.com/kata/544675c6f971f7399a000e79

func StringToNumberRun() int {
	result, _ := stringToNumber("-7")
	return result
}

func stringToNumber(s string) (int, error) {
	return strconv.Atoi(s)
}