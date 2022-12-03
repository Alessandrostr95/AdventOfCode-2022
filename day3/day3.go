package main

import (
	"fmt"
	"os"
	"strings"
)

func CommonItem(s1, s2 string) uint8 {
	// O(n^2) time
	// it is possible to compute in O(n * log(n))
	for i := range s1 {
		for j := range s2 {
			if s1[i] == s2[j] {
				return s1[i]
			}
		}
	}
	return 0
}

func CommonThree(s1, s2, s3 string) uint8 {
	for i := range s1 {
		for j := range s2 {
			for k := range s3 {
				if s1[i] == s2[j] && s1[i] == s3[k] {
					return s1[i]
				}
			}
		}
	}
	return 0
}

func main() {

	// read input file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")
	input = input[:len(input)-1]

	// --- Part One ---
	rucksacks := make([][]string, len(input))
	for i, x := range input {
		rucksacks[i] = []string{x[:len(x)/2], x[len(x)/2:]}
	}

	// find common items
	common_items := make([]uint8, len(rucksacks))
	for i, rs := range rucksacks {
		common_items[i] = CommonItem(rs[0], rs[1])
	}

	// compute priority
	priorities := make([]uint8, len(common_items))
	for i := range common_items {
		c := common_items[i]
		if c <= 122 && c >= 97 { // lowercase
			priorities[i] = c - 'a' + 1
		} else { // uppercase
			priorities[i] = c - 38
		}
	}

	// sum all priorities
	total := 0
	for _, x := range priorities {
		total += int(x)
	}

	fmt.Printf("Part 1: %d\n", total)

	// --- Part Two ---
	total = 0
	for i := 0; i < len(input); i += 3 {
		c := CommonThree(input[i], input[i+1], input[i+2])
		if c <= 122 && c >= 97 { // lowercase
			total += int(c - 'a' + 1)
		} else { // uppercase
			total += int(c - 38)
		}
	}

	fmt.Printf("Part 2: %d\n", total)

}
