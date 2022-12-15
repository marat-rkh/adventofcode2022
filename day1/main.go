package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	blocks := strings.Split(string(data), "\n\n")
	var sums []int
	for _, block := range blocks {
		entries := strings.Split(block, "\n")
		var sum int
		for _, entry := range entries {
			value, err := strconv.Atoi(entry)
			if err != nil {
				log.Fatal(err)
			}
			sum += value
		}
		sums = append(sums, sum)
	}
	sort.Ints(sums)
	top3 := sums
	if len(sums) > 3 {
		top3 = sums[len(sums)-3:]
	}
	sum := 0
	for _, value := range top3 {
		sum += value
	}
	println(sum)
}
