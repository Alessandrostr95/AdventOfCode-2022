package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Monkey data structure
type Monkey struct {
	Items     []int
	Operation func(int) int
	Test      [3]int
	Activity  int
}

/**
 * Convert a string representation of a int list to an array []int.
 * `s` must be of the form "1, 3, 5, 12, 5, 1"
 */
func ToList(s string) ([]int, error) {
	items := strings.Split(strings.ReplaceAll(s, " ", ""), ",")
	resutl := make([]int, len(items))
	for i, x := range items {
		k, err := strconv.Atoi(x)
		if err != nil {
			return resutl, err
		}
		resutl[i] = k
	}
	return resutl, nil
}

/**
 * Parses an expretion of the form `new = old [op] [num]` or `new = old [op] old`, where `op = + | *`.
 */
func ParseOperation(s string) (func(int) int, error) {
	items := strings.Split(s, " ")
	if len(items) != 5 {
		return nil, errors.New("Invalid Syntax.")
	}

	a := items[2]
	b := items[4]
	op := items[3]

	var operator func(int, int) int
	if op == "+" {
		operator = func(a, b int) int {
			return a + b
		}
	} else {
		operator = func(a, b int) int {
			return a * b
		}
	}

	var result func(int) int
	if a == b {
		result = func(old int) int {
			return operator(old, old)
		}
	} else {
		x, err := strconv.Atoi(b)
		if err != nil {
			return nil, err
		}
		result = func(old int) int {
			return operator(old, x)
		}
	}
	return result, nil
}

/**
 * Parse a test of the form:
 * Test: divisible by X
 *   If true: throw to monkey Y
 *   If false: throw to monkey Z
 *
 * Each argument `s1`, `s2` and `s3` represent a different line of the test.
 */
func ParseTest(s1, s2, s3 string) ([3]int, error) {
	r1 := strings.Split(s1, " ")
	r2 := strings.Split(s2, " ")
	r3 := strings.Split(s3, " ")
	var result [3]int

	X, err := strconv.Atoi(r1[len(r1)-1])
	if err != nil {
		return result, err
	}

	Y, err := strconv.Atoi(r2[len(r2)-1])
	if err != nil {
		return result, err
	}

	Z, err := strconv.Atoi(r3[len(r3)-1])
	if err != nil {
		return result, err
	}

	result[0] = X
	result[1] = Y
	result[2] = Z

	return result, nil
}

/**
 * Parse a Monkey given its set of rules.
 */
func ParseMonkey(s string) (*Monkey, error) {
	rows := strings.Split(s, "\n")

	if len(rows) < 6 {
		return nil, errors.New("Invalid monkey.")
	}

	items, err := ToList(strings.TrimLeft(rows[1], "Starting items: "))
	if err != nil {
		return nil, err
	}

	operation, err := ParseOperation(strings.TrimLeft(rows[2], "Operation: "))
	if err != nil {
		return nil, err
	}

	test, err := ParseTest(rows[3], rows[4], rows[5])
	if err != nil {
		return nil, err
	}

	return &Monkey{items, operation, test, 0}, nil
}

/**
 * True if a monkey `m` has no more items.
 */
func Done(m *Monkey) bool {
	return len(m.Items) == 0
}

/**
 * Execute an entire round.
 */
func Round(monkeys []*Monkey, relief func(int) int) []*Monkey {
	for _, m := range monkeys {
		for !Done(m) {
			worry, to := m.DoAction(relief)
			monkeys[to].RecieveItem(worry)
		}
	}
	return monkeys
}

/**
 * Print a list of monkey.
 */
func PrintMonkys(monkeys []*Monkey) {
	for i := range monkeys {
		fmt.Printf("Monkey %d: %v\n", i, monkeys[i].Items)
	}
}

/**
 * Execute a monkey action, i.e.:
 * - inspects the first item.
 * - calculate the new worry level.
 * - execute the test.
 *
 * Returns the worry level of the item thrown, and the index of the target monkey.
 */
func (m *Monkey) DoAction(relief func(int) int) (worry, to int) {
	m.Activity += 1

	worry = m.Items[0]

	if len(m.Items) > 1 {
		m.Items = m.Items[1:]
	} else {
		m.Items = make([]int, 0)
	}

	worry = m.Operation(worry)
	worry = relief(worry) // apply a relief reduction

	if worry%m.Test[0] == 0 {
		to = m.Test[1]
	} else {
		to = m.Test[2]
	}
	return
}

/**
 * A monkey receive an item and appends it to its list of items.
 */
func (m *Monkey) RecieveItem(x int) {
	m.Items = append(m.Items, x)
}

func main() {

	// read input file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n\n")

	// parse the monkeys
	monkeys := make([]*Monkey, len(input))
	for i, x := range input {
		m, err := ParseMonkey(x)
		if err != nil {
			panic(err)
		}
		monkeys[i] = m
	}

	// execute 20 rounds
	for i := 0; i < 20; i++ {
		monkeys = Round(monkeys, func(x int) int { return x / 3 })
	}

	// a function that compute the result.
	result := func() int64 {
		sort.SliceStable(monkeys, func(i, j int) bool {
			return monkeys[i].Activity > monkeys[j].Activity
		})
		return int64(monkeys[0].Activity) * int64(monkeys[1].Activity)
	}

	fmt.Printf("Part 1: %d\n", result())

	// --- Part Two ---

	// reset initial state
	// and compute the lcm of all tests
	lcm := 1
	for i, x := range input {
		m, err := ParseMonkey(x)
		if err != nil {
			panic(err)
		}
		monkeys[i] = m
		lcm *= m.Test[0]
	}

	// execute 10000 rounds
	for i := 1; i <= 10_000; i++ {
		k := math.Log10(float64(i))
		k = k + 1
		monkeys = Round(monkeys, func(x int) int { return x % lcm })
	}

	fmt.Printf("Part 2: %d\n", result())
}
