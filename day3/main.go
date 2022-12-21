package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	part2()
}

func part1() {
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

func part2() {
	data, err := os.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rucksacks := strings.Split(string(data), "\n")
	var badges []rune
	for _, group := range chunks(rucksacks, 3) {
		set1 := map[rune]bool{}
		for _, item1 := range group[0] {
			set1[item1] = true
		}
		common12 := map[rune]bool{}
		for _, item2 := range group[1] {
			if set1[item2] {
				common12[item2] = true
			}
		}
		for _, item3 := range group[2] {
			if common12[item3] {
				badges = append(badges, item3)
				break
			}
		}
	}
	var sum int32
	for _, badge := range badges {
		if badge < 'a' {
			sum += badge - 38
		} else {
			sum += badge - 96
		}
	}
	fmt.Println(sum)
}

func chunks(slice []string, chunkSize int) [][]string {
	var res [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		res = append(res, slice[i:end])
	}
	return res
}
