package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rounds := strings.Split(string(data), "\n")
	totalScore := 0
	for _, round := range rounds {
		letters := strings.Split(round, " ")
		opponentLetter := letters[0]
		myLetter := letters[1]
		roundScore := letterScore(myLetter)
		if isWin(opponentLetter, myLetter) {
			roundScore += 6
		} else if isDraw(opponentLetter, myLetter) {
			roundScore += 3
		}
		totalScore += roundScore
	}
	fmt.Println(totalScore)
}

func letterScore(letter string) int {
	switch letter {
	case "X":
		return 1
	case "Y":
		return 2
	default:
		return 3
	}
}

func isWin(opponentLetter string, myLetter string) bool {
	switch opponentLetter {
	case "A":
		return myLetter == "Y"
	case "B":
		return myLetter == "Z"
	default:
		return myLetter == "X"
	}
}

func isDraw(opponentLetter string, myLetter string) bool {
	switch opponentLetter {
	case "A":
		return myLetter == "X"
	case "B":
		return myLetter == "Y"
	default:
		return myLetter == "Z"
	}
}
