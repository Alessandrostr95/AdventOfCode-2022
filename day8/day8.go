package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsVisible(M [][]int, i, j int) bool {
	n, m := len(M), len(M[0])

	// check from left edge
	isVisible := true
	for jj := 0; jj < j; jj++ {
		if M[i][jj] >= M[i][j] {
			isVisible = false
		}
	}

	if isVisible {
		return isVisible
	}

	// check from right edge
	isVisible = true
	for jj := j + 1; jj < m; jj++ {
		if M[i][jj] >= M[i][j] {
			isVisible = false
		}
	}

	if isVisible {
		return isVisible
	}

	// check from top edge
	isVisible = true
	for ii := 0; ii < i; ii++ {
		if M[ii][j] >= M[i][j] {
			isVisible = false
		}
	}

	if isVisible {
		return isVisible
	}

	// check from down edge
	isVisible = true
	for ii := i + 1; ii < n; ii++ {
		if M[ii][j] >= M[i][j] {
			isVisible = false
		}
	}

	if isVisible {
		return isVisible
	}

	return false
}

func ScenicScore(M [][]int, i, j int) int {
	n, m := len(M), len(M[0])

	if i == 0 || i == n || j == 0 || j == m {
		return 0
	}

	// see left
	left := 0
	for jj := j - 1; jj >= 0; jj-- {
		left += 1
		if M[i][jj] >= M[i][j] {
			break
		}
	}

	// see right
	right := 0
	for jj := j + 1; jj < m; jj++ {
		right += 1
		if M[i][jj] >= M[i][j] {
			break
		}
	}

	// see up
	up := 0
	for ii := i - 1; ii >= 0; ii-- {
		up += 1
		if M[ii][j] >= M[i][j] {
			break
		}
	}

	// see dwon
	down := 0
	for ii := i + 1; ii < n; ii++ {
		down += 1
		if M[ii][j] >= M[i][j] {
			break
		}
	}

	return left * right * up * down
}

func main() {

	// read input file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")

	// parse matrix
	matrix := make([][]int, len(input)-1)

	for i := range matrix {
		r := strings.Split(input[i], "")
		matrix[i] = make([]int, len(r))
		for j, x := range r {
			matrix[i][j], err = strconv.Atoi(x)
			if err != nil {
				panic(err)
			}
		}
	}

	// --- Part One ---
	result := 0
	for i := range matrix {
		for j := range matrix[i] {
			if IsVisible(matrix, i, j) {
				result += 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", result)

	// --- Part Two ---

	// conpute scenic scores
	scores := make([][]int, len(matrix))
	for i := range matrix {
		scores[i] = make([]int, len(matrix[i]))
		for j := range matrix[i] {
			scores[i][j] = ScenicScore(matrix, i, j)
		}
	}

	// find max scenic score
	result = 0
	for i := range scores {
		for j := range scores[i] {
			if scores[i][j] > result {
				result = scores[i][j]
			}
		}
	}
	fmt.Printf("Part 2: %d\n", result)
}
