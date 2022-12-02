package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// read input file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")
	input = input[:len(input)-1]

	// --- Part One ---
	scores := make([]int, len(input))
	point := map[uint8]int{
		'X': 1,
		'Y': 2,
		'Z': 3,
	}
	wins_over := map[uint8]uint8{
		'X': 'C',
		'Y': 'A',
		'Z': 'B',
	}
	equiv := map[uint8]uint8{
		'X': 'A',
		'Y': 'B',
		'Z': 'C',
	}

	for i, x := range input {
		scores[i] = point[x[2]]
		if x[0] == equiv[x[2]] { // draw
			scores[i] = scores[i] + 3
		} else if x[0] == wins_over[x[2]] { // win
			scores[i] = scores[i] + 6
		}
	}

	total_score := 0
	for _, x := range scores {
		total_score += x
	}
	fmt.Printf("Part 1: %d\n", total_score)

	// --- Part Two ---
	strategy := map[string]map[uint8]uint8{
		"lose": {
			'A': 'Z',
			'B': 'X',
			'C': 'Y',
		},
		"draw": {
			'A': 'X',
			'B': 'Y',
			'C': 'Z',
		},
		"win": {
			'A': 'Y',
			'B': 'Z',
			'C': 'X',
		},
	}
	for i, x := range input {
		if x[2] == 'X' { // lose
			shape := strategy["lose"][x[0]]
			scores[i] = point[shape]
		} else if x[2] == 'Y' { // draw
			shape := strategy["draw"][x[0]]
			scores[i] = 3 + point[shape]
		} else if x[2] == 'Z' { // win
			shape := strategy["win"][x[0]]
			scores[i] = 6 + point[shape]
		}
	}

	total_score = 0
	for _, x := range scores {
		total_score += x
	}
	fmt.Printf("Part 2: %d\n", total_score)
}
