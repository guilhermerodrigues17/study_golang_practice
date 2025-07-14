package main

//Codewars -> https://www.codewars.com/kata/5672a98bdbdd995fad00000f

func RockPaperScissorsRun() string {
	return rockPaperScissors("rock", "scissors")
}

func rockPaperScissors(p1, p2  string) string {

	if p1 == p2 {
		return "Draw!"
	}

	var m = map[string]string{"rock": "paper", "paper": "scissors", "scissors": "rock"}

	if m[p1] == p2 {
		return "Player 2 won!"
	}

	return "Player 1 won!"
}