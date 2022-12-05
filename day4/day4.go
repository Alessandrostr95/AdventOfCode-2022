package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	left, right int
}

func decode(s string) (*Range, error) {
	borders := strings.Split(s, "-")
	a, err := strconv.Atoi(borders[0])
	if err != nil {
		return nil, err
	}

	b, err := strconv.Atoi(borders[1])
	if err != nil {
		return nil, err
	}

	return &Range{a, b}, nil
}

func FullyContain(r1, r2 *Range) bool {
	return r1.left <= r2.left && r1.right >= r2.right
}

func Overlaps(r1, r2 *Range) bool {
	return (r1.right >= r2.left && r1.left <= r2.left) || (r2.right >= r1.left && r2.left <= r1.left)
}

func main() {
	// read input file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(data), "\n")

	input := make([][]*Range, len(rows)-1)
	for i, row := range rows[:len(rows)-1] {
		pair := strings.Split(row, ",")

		r1, err := decode(pair[0])
		if err != nil {
			panic(err)
		}

		r2, err := decode(pair[1])
		if err != nil {
			panic(err)
		}

		input[i] = []*Range{r1, r2}
	}

	// --- Part One ---
	count := 0
	for _, pair := range input {
		if FullyContain(pair[0], pair[1]) || FullyContain(pair[1], pair[0]) {
			count += 1
		}
	}

	fmt.Printf("Part 1: %d\n", count)

	// --- Part Two ---
	count = 0
	for _, pair := range input {
		if Overlaps(pair[0], pair[1]) {
			count += 1
		}
	}
	fmt.Printf("Part 2: %d\n", count)
}
