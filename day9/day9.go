package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Rope struct
type Rope struct {
	Head []int // the haed's coordinate
	Tail *Rope // the tail
}

/**
 * Method that recursively move a rope in direction `to`.
 * I could make the code cleaner, but i don't feel like it today.
 */
func (R *Rope) MoveTo(to string) {
	H := (*R).Head
	T := (*R).Tail
	const X int = 0
	const Y int = 1

	switch to {
	case "R": // right case
		H[X] += 1

		if T != nil && H[X]-T.Head[X] > 1 {
			if H[Y] > T.Head[Y] {
				T.MoveTo("UR")
			} else if H[Y] < T.Head[Y] {
				T.MoveTo("DR")
			} else {
				T.MoveTo("R")
			}
		}

	case "U": // up case
		H[Y] += 1

		if T != nil && H[Y]-T.Head[Y] > 1 {
			if H[X] > T.Head[X] {
				T.MoveTo("UR")
			} else if H[X] < T.Head[X] {
				T.MoveTo("UL")
			} else {
				T.MoveTo("U")
			}
		}

	case "L": // left case
		H[X] -= 1

		if T != nil && H[X]-T.Head[X] < -1 {
			if H[Y] > T.Head[Y] {
				T.MoveTo("UL")
			} else if H[Y] < T.Head[Y] {
				T.MoveTo("DL")
			} else {
				T.MoveTo("L")
			}
		}

	case "D": // down case
		H[Y] -= 1

		if T != nil && H[Y]-T.Head[Y] < -1 {
			if H[X] > T.Head[X] {
				T.MoveTo("DR")
			} else if H[X] < T.Head[X] {
				T.MoveTo("DL")
			} else {
				T.MoveTo("D")
			}
		}

	case "UR":

		H[X] += 1
		H[Y] += 1

		if T != nil {
			if H[Y] == T.Head[Y]+2 {
				if H[X] > T.Head[X] {
					/* CASE
					 * | . | H |    | . | H |
					 * | . | . | -> | . | T |
					 * | T | . |    | . | . |
					 * OR
					 * | . | . | H |    | . | . | H |
					 * | . | . | . | -> | . | T | . |
					 * | T | . | . |    | . | . | . |
					 */
					T.MoveTo("UR")
				} else if H[X] == T.Head[X] {
					/* CASE
					 * | H |    | H |
					 * | . | -> | T |
					 * | T |    | . |
					 */
					T.MoveTo("U")
				}
			} else if H[X] == T.Head[X]+2 {
				if H[Y] == T.Head[Y]+1 {
					/* CASE
					 * | . | . | H | -> | . | T | H |
					 * | T | . | . |    | . | . | . |
					 */
					T.MoveTo("UR")
				} else if H[Y] == T.Head[Y] {
					/* CASE
					 * | T | . | H | -> | . | T | H |
					 */
					T.MoveTo("R")
				}
			}
		}

	case "UL":

		H[X] -= 1
		H[Y] += 1

		if T != nil {
			if H[Y] == T.Head[Y]+2 {
				if H[X] < T.Head[X] {
					/* CASE
					 * | H | . |    | H | . |
					 * | . | . | -> | T | . |
					 * | . | T |    | . | . |
					 * OR
					 * | H | . | . |    | H | . | . |
					 * | . | . | . | -> | . | T | . |
					 * | . | . | T |    | . | . | . |
					 */
					T.MoveTo("UL")
				} else if H[X] == T.Head[X] {
					/* CASE
					 * | H |    | H |
					 * | . | -> | T |
					 * | T |    | . |
					 */
					T.MoveTo("U")
				}
			} else if H[X] == T.Head[X]-2 {
				if H[Y] == T.Head[Y]+1 {
					/* CASE
					 * | H | . | . | -> | H | T | . |
					 * | . | . | T |    | . | . | . |
					 */
					T.MoveTo("UL")
				} else if H[Y] == T.Head[Y] {
					/* CASE
					 * | H | . | T | -> | H | T | . |
					 */
					T.MoveTo("L")
				}
			}
		}

	case "DR":

		H[X] += 1
		H[Y] -= 1

		if T != nil {
			if H[Y] == T.Head[Y]-2 {
				if H[X] > T.Head[X] {
					/* CASE
					 * | T | . |    | . | . |
					 * | . | . | -> | . | T |
					 * | . | H |    | . | H |
					 * OR
					 * | T | . | . |    | . | . | . |
					 * | . | . | . | -> | . | T | . |
					 * | . | . | H |    | . | . | H |
					 */
					T.MoveTo("DR")
				} else if H[X] == T.Head[X] {
					/* CASE
					 * | T |    | . |
					 * | . | -> | T |
					 * | H |    | H |
					 */
					T.MoveTo("D")
				}
			} else if H[X] == T.Head[X]+2 {
				if H[Y] == T.Head[Y]-1 {
					/* CASE
					 * | T | . | . | -> | . | . | . |
					 * | . | . | H |    | . | T | H |
					 */
					T.MoveTo("DR")
				} else if H[Y] == T.Head[Y] {
					/* CASE
					 * | T | . | H | -> | . | T | H |
					 */
					T.MoveTo("R")
				}
			}
		}

	case "DL":

		H[X] -= 1
		H[Y] -= 1

		if T != nil {
			if H[Y] == T.Head[Y]-2 {
				if H[X] < T.Head[X] {
					/* CASE
					 * | . | T |    | . | . |
					 * | . | . | -> | T | . |
					 * | H | . |    | H | . |
					 * OR
					 * | . | . | T |    | . | . | . |
					 * | . | . | . | -> | . | T | . |
					 * | H | . | . |    | H | . | . |
					 */
					T.MoveTo("DL")
				} else if H[X] == T.Head[X] {
					/* CASE
					 * | T |    | . |
					 * | . | -> | T |
					 * | H |    | H |
					 */
					T.MoveTo("D")
				}
			} else if H[X] == T.Head[X]-2 {
				if H[Y] == T.Head[Y]-1 {
					/* CASE
					 * | . | . | T | -> | . | . | . |
					 * | H | . | . |    | H | T | . |
					 */
					T.MoveTo("DL")
				} else if H[Y] == T.Head[Y] {
					/* CASE
					 * | H | . | T | -> | H | T | . |
					 */
					T.MoveTo("L")
				}
			}
		}
	}
}

/**
 * Print the position of all `rope`'s node in a map `H`x`W`, setting the origin in pos (dx, dy).
 */
func printMap(W, H, dx, dy int, rope []*Rope) {
	for i := H - 1; i >= 0; i-- {
		for j := 0; j < W; j++ {
			cell := ". "
			for k := len(rope) - 1; k >= 0; k-- {
				r := rope[k]
				if r.Head[0]+dx == j && r.Head[1]+dy == i {
					cell = fmt.Sprintf("%d ", k)
				}
			}
			fmt.Print(cell)
		}
		fmt.Print("\n")
	}
}

func main() {

	// read input file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")

	// --- Part One ---
	// NOTE: part one was written **before** the needed of `Rope` structure.

	H := []int{0, 0} // (x,y) coordinate starting from BOTTOME-LEFT ORIGIN
	T := []int{0, 0}

	const X int = 0
	const Y int = 1

	// return true if two point ar adjacent
	isNear := func() bool {
		return !(math.Abs(float64(H[X]-T[X])) > 1 || math.Abs(float64(H[Y]-T[Y])) > 1)
	}

	move := func(to string) {
		switch to {
		case "R": // right case
			H[X] += 1
			if !isNear() {
				T[X] = H[X] - 1
				T[Y] = H[Y]
			}
		case "U": // up case
			H[Y] += 1
			if !isNear() {
				T[X] = H[X]
				T[Y] = H[Y] - 1
			}
		case "L": // left case
			H[X] -= 1
			if !isNear() {
				T[X] = H[X] + 1
				T[Y] = H[Y]
			}
		case "D":
			H[Y] -= 1
			if !isNear() {
				T[X] = H[X]
				T[Y] = H[Y] + 1
			}
		}
	}

	type Pair struct {
		x, y int
	}

	count := make(map[Pair]bool)
	for i := 0; i < len(input)-1; i++ {
		action := strings.Split(input[i], " ")

		to := action[0]

		k, err := strconv.Atoi(action[1])
		if err != nil {
			panic(err)
		}

		for j := 0; j < k; j++ {
			move(to)
			count[Pair{T[X], T[Y]}] = true
		}
	}

	fmt.Printf("Part 1: %d\n", len(count))

	// --- Part Two ---

	// create the rope's knots
	rope := make([]*Rope, 10)
	for i := range rope {
		rope[i] = &Rope{[]int{0, 0}, nil}
	}

	// connect al knots
	for i := 0; i < len(rope)-1; i++ {
		rope[i].Tail = rope[i+1]
	}

	count = make(map[Pair]bool)
	for i := 0; i < len(input)-1; i++ {
		action := strings.Split(input[i], " ")

		to := action[0]

		k, err := strconv.Atoi(action[1])
		if err != nil {
			panic(err)
		}

		// fmt.Printf("== %s ==\n", input[i])
		for j := 0; j < k; j++ {
			rope[0].MoveTo(to)
			tail := rope[9]
			count[Pair{tail.Head[X], tail.Head[Y]}] = true
		}
		// printMap()
		// fmt.Print("\n--------\n")
	}

	fmt.Printf("Part 2: %d\n", len(count))
}
