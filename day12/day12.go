package main

import (
	"fmt"
	"os"
	"strings"
)

const UP = 0
const LEFT = 1
const RIGHT = 2
const DOWN = 3

/**
 * Utility funciton for parsing input.
 */
func ToMatrix(rows []string) [][]uint8 {
	result := make([][]uint8, len(rows))
	for i, r := range rows {
		result[i] = make([]uint8, len(r))
		for j := range r {
			result[i][j] = r[j]
		}
	}
	return result
}

/**
 * Printss the shortest path from `s` to `t` on the `heatmap`.
 */
func PrintPath(heatmap [][]uint8, s, t [2]int) {
	SPT := BFS(heatmap, s)

	M := make([][]uint8, len(heatmap))
	for i := range heatmap {
		M[i] = make([]uint8, len(heatmap[i]))
		for j := range heatmap[i] {
			M[i][j] = heatmap[i][j]
		}
	}
	M[t[0]][t[1]] = 'E'

	current := t
	for SPT[current] != current {
		pred := SPT[current]
		if pred[0] == current[0]-1 { // from top
			M[pred[0]][pred[1]] = 'V'
		} else if pred[0] == current[0]+1 { // from bottom
			M[pred[0]][pred[1]] = '^'
		} else if pred[1] == current[1]-1 { // from left
			M[pred[0]][pred[1]] = '>'
		} else { // from right
			M[pred[0]][pred[1]] = '<'
		}
		current = pred
	}

	for _, row := range M {
		for _, c := range row {
			s := EmbedChar(uint8(c))
			fmt.Printf(s)
		}
		fmt.Println()
	}
}

/**
 * Given a cell (`i`,`j`) on a matrix `n` x `m`, return the direction of its adiacent neighbors.
 */
func Neighborhood(i, j, n, m int) (neigh []int) {

	if i == 0 && j == 0 { // top-left edge
		neigh = []int{
			RIGHT,
			DOWN,
		}
		return
	}

	if i == 0 && j == m-1 { // top-right edge
		neigh = []int{
			LEFT,
			DOWN,
		}
		return
	}

	if i == 0 && j > 0 && j < m-1 { // top edge
		neigh = []int{
			LEFT,
			DOWN,
			RIGHT,
		}
		return
	}

	if i == n-1 && j == 0 { // bottom-left edge
		neigh = []int{
			UP,
			RIGHT,
		}
		return
	}

	if i == n-1 && j == m-1 { // bottom-right edge
		neigh = []int{
			UP,
			LEFT,
		}
		return
	}

	if i == n-1 && j > 0 && j < m-1 { // bottom-middle edge
		neigh = []int{
			LEFT,
			UP,
			RIGHT,
		}
		return
	}

	if j == 0 && i > 0 && i < n-1 { // left-middle edge
		neigh = []int{
			UP,
			RIGHT,
			DOWN,
		}
		return
	}

	if j == m-1 && i > 0 && i < n-1 { // right-middle edge
		neigh = []int{
			UP,
			LEFT,
			DOWN,
		}
		return
	}

	// middle cell
	neigh = []int{
		UP,
		LEFT,
		RIGHT,
		DOWN,
	}
	return
}

/**
 * Returns the BFS-tree of the `heatmap` rooted in `S`.
 */
func BFS(heatmap [][]uint8, e [2]int) (predecessor map[[2]int][2]int) {
	// (inner)-SPT
	predecessor = make(map[[2]int][2]int)

	// map of seen nodes
	seen := make([][]bool, len(heatmap))

	for i := range heatmap {
		seen[i] = make([]bool, len(heatmap[i]))
		for j := range heatmap[i] {
			seen[i][j] = false
		}
	}

	queue := [][2]int{e}
	seen[e[0]][e[1]] = true
	predecessor[e] = e
	n := len(heatmap)
	m := len(heatmap[0])

	for len(queue) > 0 {
		u := queue[0]
		if len(queue) == 1 {
			queue = make([][2]int, 0)
		} else {
			queue = queue[1:]
		}

		i := u[0]
		j := u[1]

		for _, v := range Neighborhood(i, j, n, m) {
			ii := i
			jj := j

			switch v {
			case UP:
				ii -= 1
			case LEFT:
				jj -= 1
			case RIGHT:
				jj += 1
			case DOWN:
				ii += 1
			}

			if !seen[ii][jj] && heatmap[ii][jj] <= heatmap[i][j]+1 {
				seen[ii][jj] = true
				predecessor[[2]int{ii, jj}] = u
				queue = append(queue, [2]int{ii, jj})
			}
		}
	}

	return
}

/**
 * Returns the distances from `s` to `t` on the `heatmap`.
 * If the second return values is `false` it means that `t` is not reachable from `s`.
 */
func Dist(heatmap [][]uint8, s, t [2]int) (int, bool) {
	SPT := BFS(heatmap, s)
	dist := 0
	x := t
	for SPT[x] != x {
		if _, ok := SPT[x]; !ok {
			return -1, false
		}
		dist += 1
		x = SPT[x]
	}
	return dist, true
}

func main() {

	// read input file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")

	// -- Part One --
	heatmap := ToMatrix(input[:len(input)-1])
	var S [2]int // the index of the starting point
	var E [2]int // the index of the target point
	for i := range heatmap {
		for j := range heatmap[i] {
			if heatmap[i][j] == 'S' {
				heatmap[i][j] = 'a'
				S = [2]int{i, j}
			} else if heatmap[i][j] == 'E' {
				heatmap[i][j] = 'z' + 1
				E = [2]int{i, j}
			}
		}
	}

	result, _ := Dist(heatmap, S, E)
	fmt.Printf("Part One: %d\n", result)
	PrintPath(heatmap, S, E) // comment here

	// --- Part Two ---
	X := S
	for i := range heatmap {
		for j := range heatmap[i] {

			if heatmap[i][j] == 'a' {
				k, ok := Dist(heatmap, [2]int{i, j}, E)
				if ok && k < result {
					X = [2]int{i, j}
					result = k
				}
			}

		}
	}
	fmt.Printf("Part Two: %d\n", result)
	PrintPath(heatmap, X, E) // comment here
}
