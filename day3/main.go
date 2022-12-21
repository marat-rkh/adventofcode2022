package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rucksacks := strings.Split(string(data), "\n")
	var mistakes []rune
	for _, rucksack := range rucksacks {
		compartment1 := rucksack[:len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2:]
		set1 := map[rune]bool{}
		for _, item1 := range compartment1 {
			set1[item1] = true
		}
		for _, item2 := range compartment2 {
			if set1[item2] {
				mistakes = append(mistakes, item2)
				break
			}
		}
	}
	var sum int32
	for _, mistake := range mistakes {
		if mistake < 'a' {
			sum += mistake - 38
		} else {
			sum += mistake - 96
		}
	}
	fmt.Println(sum)
}
