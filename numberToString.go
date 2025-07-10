package main

import "strconv"

func NumberToStringRun() string {
	return numberToString(10)
}

func numberToString(n int) string {
	return strconv.Itoa(n)
}