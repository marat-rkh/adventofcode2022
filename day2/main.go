package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var letterScore = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var win = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var draw = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var loose = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

var roundResult = map[string]map[string]string{
	"X": loose,
	"Y": draw,
	"Z": win,
}

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
		myLetter := roundResult[letters[1]][opponentLetter]
		roundScore := letterScore[myLetter]
		if win[opponentLetter] == myLetter {
			roundScore += 6
		} else if draw[opponentLetter] == myLetter {
			roundScore += 3
		}
		totalScore += roundScore
	}
	fmt.Println(totalScore)
}
