package main

import (
	"log"
	"strconv"
	"strings"

	utils "github.com/ssedrick/adventofcode2021"
)

func main() {
	// lines, err := utils.ReadLines("input.test.txt")
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		log.Fatalf("Error reading lines: %s", err)
	}

	horizontal, depth, aim := 0, 0, 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		mag, err := strconv.Atoi(parts[1])

		if err != nil {
			log.Fatalf("Error parsing int: %s", err)
		}

		log.Println(horizontal, depth, aim)
		log.Println(direction, mag)

		switch direction {
		case "up":
			aim = aim - mag
		case "down":
			aim = aim + mag
		case "forward":
			horizontal = horizontal + mag
			depth = depth + aim*mag
		}
	}

	log.Printf("Depth: %d, Horizontal: %d, answer: %d", depth, horizontal, horizontal*depth)
}
