package main

import (
	"fmt"
	"os"
)

/**
 * Returns true if in the subsequence of `s` starting from `from` and with length `k` contains only different chars.
 * Time O(k\log{k})
 */
func KDistincts(s string, from int, k int) bool {
	seen := make(map[uint8]bool, k)
	for i := from; i < from+k; i++ {
		seen[s[i]] = false
	}
	for i := from; i < from+k; i++ {
		if !seen[s[i]] {
			seen[s[i]] = true
		} else {
			return false
		}
	}
	return true
}

func main() {

	// read input file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := string(data[:len(data)-1])

	// --- Part One ---
	for i := 0; i < len(input)-4; i++ {
		a, b, c, d := input[i], input[i+1], input[i+2], input[i+3]
		if a != b && a != c && a != d && b != c && b != d && c != d {
			fmt.Printf("Part 1: %d\n", i+4)
			break
		}
	}

	// --- Part Two ---
	k := 14
	for i := 0; i < len(input)-k; i++ {
		if KDistincts(input, i, k) {
			fmt.Printf("Part 2: %d\n", i+k)
			break
		}
	}
}
