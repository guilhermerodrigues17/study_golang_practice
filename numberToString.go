package main

import "strconv"

//Codewars -> https://www.codewars.com/kata/5265326f5fda8eb1160004c8

func NumberToStringRun() string {
	return numberToString(10)
}

func numberToString(n int) string {
	return strconv.Itoa(n)
}