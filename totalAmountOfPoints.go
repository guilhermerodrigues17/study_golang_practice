package main

//Codewars -> https://www.codewars.com/kata/5bb904724c47249b10000131

func TotalAmountOfPointsRun() int {
	return totalAmountOfPoints([]string{"0:1","0:2","0:3","0:4","1:2","1:3","1:4","2:3","2:4","3:4"})
}

func totalAmountOfPoints(games []string) int {
	var totalPoints int

	for _, game := range games {
		xPoints, yPoints := game[0], game[2]
		
		if xPoints == yPoints {
			totalPoints++
		}

		if xPoints > yPoints {
			totalPoints += 3
		}
	}

	return totalPoints
}

/*
[x:y] -> placar
["3:1", "2:2", "0:1", ...]

if x > y: 3 points (win)
if x < y: 0 points (loss)
if x = y: 1 point (tie)
*/