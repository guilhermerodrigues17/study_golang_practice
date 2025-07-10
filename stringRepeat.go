package main

import (
	"log"
	"strings"
)

//Codewars -> https://www.codewars.com/kata/57a0e5c372292dd76d000d7e

func StringRepeatRun() string {
	return stringRepeat(5, "Hello")
}

func stringRepeat(n int, s string) string {
	if n < 0 {
		log.Fatal("n must be a positive integer")
	}
	return strings.Repeat(s, n)
}