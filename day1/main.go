package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	blocks := strings.Split(string(data), "\n\n")
	max := 0
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
		if sum > max {
			max = sum
		}
	}
	println(max)
}
