package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

/**
 * Push element on top of the stack.
 */
func (s *Stack) Push(el string) {
	*s = append(*s, el)
}

/**
 * Returns true if the stack is empty, false otherwise.
 */
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Length() int {
	return len(*s)
}

/**
 * Pop element from top of the stack.
 * If there is no elements on the stack, returns false.
 */
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	n := s.Length() - 1
	el := (*s)[n]
	*s = (*s)[:n]
	return el, true
}

/**
 * Show the element on top of the stack.
 * If there is no elements on the stack, returns false.
 */
func (s *Stack) Head() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	n := s.Length() - 1
	return (*s)[n], true
}

/**
 * Parse the stack input string and return an array of stacks.
 */
func DecodeStacks(s string) ([]*Stack, error) {
	rows := strings.Split(s, "\n") // split all rows

	// get the number n of stacks
	last_row := rows[len(rows)-1]
	last_row_nums := strings.Split(last_row, "  ")
	n, err := strconv.Atoi(strings.Trim(last_row_nums[len(last_row_nums)-1], " "))
	if err != nil {
		return nil, err
	}

	// initialize empty stacks
	stacks := make([]*Stack, n)
	for i := range stacks {
		stacks[i] = &Stack{}
	}

	for _, r := range rows[:len(rows)-1] {
		for i := 1; i <= n*4; i += 4 {
			if r[i] != ' ' {
				*(stacks[(i-1)/4]) = append([]string{string(r[i])}, *(stacks[(i-1)/4])...)
			}
		}
	}

	return stacks, nil
}

/**
 * Parse the moves input string and return an array of []int.
 */
func DecodeMoves(s string) ([][]int, error) {
	rows := strings.Split(s, "\n")

	n := len(rows) - 1
	moves := make([][]int, n)
	for i := 0; i < n; i++ {
		x := strings.Split(rows[i], " ")

		move, err := strconv.Atoi(x[1])
		if err != nil {
			return nil, err
		}
		from, err := strconv.Atoi(x[3])
		if err != nil {
			return nil, err
		}
		to, err := strconv.Atoi(x[5])
		if err != nil {
			return nil, err
		}

		moves[i] = []int{move, from, to}
	}

	return moves, nil
}

/**
 * Execute a move like:
 * move `m[0]` elements **in sequence** from stack `m[1]` to stack `m[2]`.
 */
func move(stacks []*Stack, m []int) bool {
	for i := 0; i < m[0]; i++ {
		el, ok := stacks[m[1]-1].Pop()
		if !ok {
			return ok
		}

		stacks[m[2]-1].Push(el)
	}
	return true
}

/**
 * Execute a move like:
 * move `m[0]` elements **in block** from stack `m[1]` to stack `m[2]`.
 */
func move2(stacks []*Stack, m []int) bool {
	var s Stack
	for i := 0; i < m[0]; i++ {
		el, ok := stacks[m[1]-1].Pop()
		if !ok {
			return ok
		}
		s.Push(el)
	}

	for i := 0; i < m[0]; i++ {
		el, ok := s.Pop()
		if !ok {
			return ok
		}
		stacks[m[2]-1].Push(el)
	}
	return true
}

func main() {

	// read input file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n\n")
	// fmt.Println(input[0])

	// --- Part One ---

	// decode stacks
	stacks, err := DecodeStacks(input[0])
	if err != nil {
		panic(err)
	}

	// decode moves
	moves, err := DecodeMoves(input[1])
	if err != nil {
		panic(err)
	}

	// do moves
	for _, m := range moves {
		move(stacks, m)
	}

	result := make([]string, len(stacks))
	for i, s := range stacks {
		el, _ := s.Head()
		result[i] = el
	}
	fmt.Printf("Part 1: %s\n", strings.Join(result, ""))

	// --- Part 2 ---

	// reset stacks
	stacks, err = DecodeStacks(input[0])
	if err != nil {
		panic(err)
	}

	// do moves
	for _, m := range moves {
		move2(stacks, m)
	}

	result = make([]string, len(stacks))
	for i, s := range stacks {
		el, _ := s.Head()
		result[i] = el
	}
	fmt.Printf("Part 2: %s\n", strings.Join(result, ""))
}
