package main

import (
	"fmt"
	"strings"
)

//Codewars -> https://www.codewars.com/kata/56f3a1e899b386da78000732

func PartsListRun() string {
	return partsList([]string{"cdIw", "tzIy", "xDu", "rThG"})
}

func partsList(arr []string) string {
	var result string
	for i := 1; i < len(arr); i++ {
		tempExpression := fmt.Sprintf("(%s, %s)", strings.Join(arr[:i], " "), strings.Join(arr[i:], " "))
		result += tempExpression
	}
	return result
}

//retorna "(cdIw, tzIy xDu rThG)(cdIw tzIy, xDu rThG)(cdIw tzIy xDu, rThG)"
