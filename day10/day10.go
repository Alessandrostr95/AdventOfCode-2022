package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// read input file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")

	// --- Part One ---

	clock := 1    // the clock
	X := 1        // the X register
	IR := 0       // the instruction register
	addX := false // utility flag
	k := 0        // quantity to add to register X

	check := func() bool { // check function
		return clock == 20 || clock == 60 || clock == 100 || clock == 140 || clock == 180 || clock == 220
	}

	result := 0
	for IR < len(input)-1 {
		if check() {
			result += clock * X
		}

		clock += 1 // increment the clock

		if addX { // if it is the second clock on a `addX` instruction
			addX = false // set the flag to false
			X += k       // sum k to X
		} else {
			cmd := strings.Split(input[IR], " ") // read next instruction
			IR += 1                              // increment the IR

			if len(cmd) == 2 { // if cmd is an `addX` instruction
				k, err = strconv.Atoi(cmd[1]) // parse k
				if err != nil {
					panic(err)
				}
				addX = true // set addX flag to true
			}
		}
	}

	fmt.Printf("Part 1: %d\n", result)

	// --- Part Two ---
	clock = 1
	X = 1
	IR = 0
	addX = false
	k = 0

	fmt.Println("Part Two:")

	// a simple double for loop, with checking if CTR is near X
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {

			// check what to print
			if j >= X-1 && j <= X+1 {
				fmt.Print("\033[48;2;49;94;104m \033[0m")
			} else {
				fmt.Print(".")
			}

			// same code of part one
			clock += 1
			if addX {
				addX = false
				X += k
			} else {
				cmd := strings.Split(input[IR], " ")
				IR += 1

				if len(cmd) == 2 {
					k, err = strconv.Atoi(cmd[1])
					if err != nil {
						panic(err)
					}
					addX = true
				}
			}

		}
		fmt.Println()
	}
}
