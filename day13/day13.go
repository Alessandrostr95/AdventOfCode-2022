package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

/**
 * Splits a string s in a list of token strings composed by "[", "]" or a number string.
 */
func Tokens(s string) []string {
	s = strings.ReplaceAll(s, "]", ",]")
	s = strings.ReplaceAll(s, "[", "[,")

	allTokens := strings.Split(s, ",")
	result := make([]string, 0)
	for _, t := range allTokens {
		if !(t == "," || t == "") {
			result = append(result, t)
		}
	}
	return result
}

/**
 * Utility function for parsing packets.
 */
func RecursiveParse(tokens []string, closedBy map[int]int, from int) ([]interface{}, error) {
	_, ok := closedBy[from]
	if !ok {
		return nil, errors.New("Invalid Syntax. [#3]")
	}

	result := make([]interface{}, 0)
	if from >= closedBy[from] {
		return result, nil
	}

	for k := from + 1; k < closedBy[from]; k++ {
		t := tokens[k]
		if n, err := strconv.Atoi(t); err == nil {
			result = append(result, n)
		} else if t == "[" {
			x, err := RecursiveParse(tokens, closedBy, k)
			if err != nil {
				return nil, err
			}
			result = append(result, x)
			k = closedBy[k]
		}
	}
	return result, nil
}

/**
 * Given an input row, this function pares it.
 */
func Parse(s string) ([]interface{}, error) {
	if s == "" {
		return nil, errors.New("Invalid Syntax. [#0]")
	}
	tokens := Tokens(s)

	closedBy := make(map[int]int)
	par := &Stack{}
	for i, t := range tokens {
		if t == "[" {
			par.Push(i)
		} else if t == "]" {
			j, ok := par.Pop()
			if !ok {
				return nil, errors.New("Invalid Syntax. [#1]")
			}
			closedBy[j] = i
		}
	}

	if !par.IsEmpty() {
		return nil, errors.New("Invalid Syntax. [#2]")
	}

	result, err := RecursiveParse(tokens, closedBy, 0)
	if err != nil {
		return nil, err
	}

	return result, nil
}

const RIGHT_ORDER = 0
const WRONG_ORDER = 1
const SAME = 2

/**
 * This function compares two packets `s1` and `s2`, according the AOC day 13 rules.
 * It returns
 * - `RIGHT_ORDER` if `s1` is lower than `s2`.
 * - `WRONG_ORDER` if `s2` is lower than `s1`.
 * - `SAME` if `s1` and `s2` are the same.
 */
func Compare(s1, s2 []interface{}) int {
	n := len(s1)
	m := len(s2)
	i, j := 0, 0

	for i < n && j < m {
		left, right := s1[i], s2[j]
		i += 1
		j += 1

		// CASE 1: both values are integers.
		if reflect.TypeOf(left) == reflect.TypeOf(1) && reflect.TypeOf(right) == reflect.TypeOf(1) {
			if left.(int) < right.(int) {
				return RIGHT_ORDER
			} else if left.(int) > right.(int) {
				return WRONG_ORDER
			} else {
				continue
			}
		}

		// CASE 2: both values are lists.
		if reflect.TypeOf(left) == reflect.TypeOf(s1) && reflect.TypeOf(right) == reflect.TypeOf(s1) {
			result := Compare(left.([]interface{}), right.([]interface{}))
			if result != SAME {
				return result
			} else {
				continue
			}
		}

		// CASE 3: exactly one value is an integer.
		if reflect.TypeOf(left) == reflect.TypeOf(1) { // first element is the integer
			list := []interface{}{left.(int)}
			result := Compare(list, right.([]interface{}))
			if result != SAME {
				return result
			} else {
				continue
			}
		} else { // second element is the integer
			list := []interface{}{right.(int)}
			result := Compare(left.([]interface{}), list)
			if result != SAME {
				return result
			} else {
				continue
			}
		}
	}

	if i == n && j < m { // left list runs out of items first.
		return RIGHT_ORDER
	} else if j == m && i < n { // right list runs out of items first.
		return WRONG_ORDER
	}

	return SAME
}

func main() {

	// read input file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")

	// --- Part One ---

	// parse packets
	packets := make([][]interface{}, 0)
	for i := 0; i < len(input)-1; i++ {
		if l, err := Parse(input[i]); err == nil {
			packets = append(packets, l)
		}
	}

	result := 0
	for i := 0; i < len(packets); i += 2 {
		x := Compare(packets[i], packets[i+1])
		if x == RIGHT_ORDER { // if a pair is the right order
			result += (i + 2) / 2 // sum their (pair) index (starting from 1)
		}
	}

	fmt.Printf("Part One: %d\n", result)

	// --- Part Two ---

	// insert additional divider packets
	div1, _ := Parse("[[2]]")
	packets = append(packets, div1)

	div2, _ := Parse("[[6]]")
	packets = append(packets, div2)

	// sort packets according rules
	sort.SliceStable(packets, func(i, j int) bool {
		x := Compare(packets[i], packets[j])
		return x == RIGHT_ORDER
	})

	// compute solution
	result = 1
	for i := range packets {
		if Compare(packets[i], div1) == SAME || Compare(packets[i], div2) == SAME {
			result *= (i + 1)
		}
	}

	fmt.Printf("Part Two: %d\n", result)
}
