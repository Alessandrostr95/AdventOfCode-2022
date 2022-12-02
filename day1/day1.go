package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toList(s string) []int {
	nums := strings.Split(s, "\n")
	result := make([]int, 0)
	for _, n := range nums {
		k, err := strconv.Atoi(n)
		if err == nil { // just ignore parsing errors
			result = append(result, k)
		}
	}
	return result
}

func sum(arr []int) int {
	result := 0
	for _, n := range arr {
		result += n
	}
	return result
}

func main() {

	// --- Day 1: Calorie Counting ---

	data, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	// MAP phase
	items_calories := make([][]int, 0)
	for _, v := range strings.Split(string(data), "\n\n") {
		l := toList(v)
		items_calories = append(items_calories, l)
	}

	// REDUCE phase
	amount_calories := make([]int, 0)
	for _, l := range items_calories {
		amount_calories = append(amount_calories, sum(l))
	}

	// finding the maximum's index
	ind_max := 0
	for i := range amount_calories {
		if amount_calories[i] > amount_calories[ind_max] {
			ind_max = i
		}
	}

	fmt.Printf("Part 1: %d\n", amount_calories[ind_max])

	// --- Part Two ---

	// need to fine the second with more calories
	snd_ind_max := 0
	for i := range amount_calories {
		if amount_calories[i] > amount_calories[snd_ind_max] && i != ind_max {
			snd_ind_max = i
		}
	}

	// now the last (third) one
	trd_ind_max := 0
	for i := range amount_calories {
		if amount_calories[i] > amount_calories[trd_ind_max] && i != ind_max && i != snd_ind_max {
			trd_ind_max = i
		}
	}

	result := amount_calories[ind_max] + amount_calories[snd_ind_max] + amount_calories[trd_ind_max]
	fmt.Printf("Part 2: %d\n", result)
}
