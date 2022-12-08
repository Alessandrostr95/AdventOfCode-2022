package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * File System Node
 */
type FSNode struct {
	IsDir    bool      // true if the node is a directory
	Name     string    // the name of the node
	Father   *FSNode   // the father node
	Children []*FSNode // Children nodes
	Size     int       // the size of the node
}

/**
 * Compute recursively the size of all direcotry
 */
func (v *FSNode) RecursiveComputeSize() int {
	if v == nil { // just in case
		return 0
	}

	if !(*v).IsDir {
		return (*v).Size
	}

	for _, child := range (*v).Children {
		(*v).Size += child.RecursiveComputeSize()
	}

	return (*v).Size
}

/**
 * Parse an input line to FSNode
 */
func ParseNode(line string) (*FSNode, error) {
	field := strings.Split(line, " ")
	if field[0] == "dir" {
		dirNode := &FSNode{true, field[1], currentDir, []*FSNode{}, 0}
		return dirNode, nil
	}
	size, err := strconv.Atoi(field[0])
	if err != nil {
		return nil, err
	}
	fileNode := &FSNode{false, field[1], currentDir, nil, size}
	return fileNode, nil
}

func HasChild(father *FSNode, name string) bool {
	for _, x := range (*father).Children {
		if (*x).Name == name {
			return true
		}
	}
	return false
}

var root *FSNode = &FSNode{true, "/", nil, []*FSNode{}, 0} // root direcotry
var currentDir *FSNode = root                              // current directory

/**
 * Sets the current direcotry
 */
func Move(dir string) bool {
	switch dir {
	case "..":
		currentDir = (*&currentDir).Father
		return true
	case "/":
		currentDir = root
		return true
	default:
		for _, child := range (*currentDir).Children {
			if (*child).IsDir && (*child).Name == dir {
				currentDir = child
				return true
			}
		}
	}
	return false // do nothing
}

func main() {

	// read input file
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")

	// --- Part One ---

	// parsing file sistem tree
	for i := 0; i < len(input)-1; i++ {
		line := strings.Split(input[i], " ")

		if line[0] == "$" {
			if line[1] == "cd" { // move command
				Move(line[2])
			}
			// don't need ls command
		} else {

			// if a file or dir is not already seen
			if !HasChild(currentDir, line[1]) {
				node, err := ParseNode(input[i])
				if err != nil {
					panic(err)
				}
				currentDir.Children = append(currentDir.Children, node)
			}

		}
	}

	// compute all sizes
	root.RecursiveComputeSize()

	// bfs traversal
	queue := []*FSNode{root}
	result := 0

	for len(queue) != 0 {
		// dequeue the last element of the queue
		u := queue[0]
		if len(queue) == 1 {
			queue = []*FSNode{}
		} else {
			queue = queue[1:]
		}

		if (*u).IsDir && (*u).Size <= 100000 {
			result += (*u).Size
		}

		// enqueue u's children
		for _, v := range (*u).Children {
			queue = append(queue, v)
		}
	}

	fmt.Printf("Day 1: %d\n", result)

	// --- Part Two ---

	// unused_space
	unused_space := 70_000_000 - (*root).Size

	// a list of all candidate result
	candidate_size := make([]int, 0)

	// bfs search to find all possible candidate solution
	queue = []*FSNode{root}

	for len(queue) != 0 {
		// dequeue the last element of the queue
		u := queue[0]
		if len(queue) == 1 {
			queue = []*FSNode{}
		} else {
			queue = queue[1:]
		}

		if (*u).IsDir && (*u).Size+unused_space >= 30_000_000 {
			candidate_size = append(candidate_size, (*u).Size)
		}

		// enqueue u's children
		for _, v := range (*u).Children {
			queue = append(queue, v)
		}
	}

	// find solution, i.e. the minimum candidate size
	result = candidate_size[0]
	for _, s := range candidate_size {
		if s < result {
			result = s
		}
	}

	fmt.Printf("Day 2: %d\n", result)
}
