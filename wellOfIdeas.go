package main

//Codewars -> https://www.codewars.com/kata/57f222ce69e09c3630000212

func WellOfIdeasRun() string {

	return wellOfIdeas([]string{"bad", "bad", "bad", "bad", "good", "bad", "good", "good"})
}

func wellOfIdeas(x []string) string {
	var count int
	
	for _, idea := range x {
		if idea == "good" {
			count++
		}
	}

	if count == 0 {	
		return "Fail!"
	}

	if count > 2 {
		return "I smell a series!"
	}

	return "Publish!"

}