package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// coordinate data structure.
type Coord [2]int

/**
 * Gets the x-value of a coordinate.
 */
func (c *Coord) X() int {
  return (*c)[0]
}

/**
 * Gets the y-value of a coordinate.
 */
func (c *Coord) Y() int {
  return (*c)[1]
}

/**
 * Stes the x-value of a coordinate.
 */
func (c *Coord) SetX(x int) {
  (*c)[0] = x
}

/**
 * Stes the y-value of a coordinate.
 */
func (c *Coord) SetY(y int) {
  (*c)[1] = y
}

/**
 * Moves down a coordinate.
 */
func (c *Coord) Down() *Coord {
  return &Coord{c.X(), c.Y()+1}
}

/**
 * Moves down-left a coordinate.
 */
func (c *Coord) DownLeft() *Coord {
  return &Coord{c.X()-1, c.Y()+1}
}

/**
 * Moves down-right a coordinate.
 */
func (c *Coord) DownRight() *Coord {
  return &Coord{c.X()+1, c.Y()+1}
}

/**
 * Parses an input row into a list of coordantes.
 */
func ParseRow(s string) ([]*Coord, error) {
  
  result := make([]*Coord, 0)

  for _, e := range strings.Split(s, " -> ") {
    p := strings.Split(e, ",")
    
    x, err := strconv.Atoi(p[0])
    if err != nil {
      return nil, err
    }

    y, err  := strconv.Atoi(p[1])
    if err != nil {
      return nil, err
    }
    result = append(result, &Coord{x,y})
  }
  
  return result, nil
}

/**
 * Checks if a coordinate `p` is on the line passing to `a` and `b`.
 */
func IntersectLine(a, b, p *Coord) bool {

  if a.X() == b.X() && p.X() == a.X() {
    if (a.Y() <= p.Y() && p.Y() <= b.Y()) || (b.Y() <= p.Y() && p.Y() <= a.Y()) {
      return true
    }
  }

  if a.Y() == b.Y() && p.Y() == a.Y() {
    if (a.X() <= p.X() && p.X() <= b.X()) || (b.X() <= p.X() && p.X() <= a.X()) {
      return true
    }
  }

  return false
}

/**
 * Generate a map with keys = coordante and value = wall
 */
func GenerateMap(rows, cols int, walls [][]*Coord) map[Coord]int8 {
  result := make(map[Coord]int8)

  for i := 0; i <= rows; i++ {
    for j := 0; j <= cols; j++ {
      p := &Coord{j, i}
      
      isWall := false
      for _, w := range walls {
        for k := 0; k < len(w)-1; k++ {
          isWall = isWall || IntersectLine(w[k], w[k+1], p)
        }
      }

      if isWall {
        result[*p] = '#' 
      }
    }
  }
    
  return result
}

/**
 * Print the current state of the map on the standard output.
 */
func PrintMap(matrix map[Coord]int8, minX, minY int, maxX, maxY int) {

  // ┌┘└┐─│
  fmt.Print("┌")
  for x := minX; x <= maxX ; x++ {
    fmt.Print("─")
  }
  fmt.Println("┐")

  for y := minY; y <= maxY; y++ {
    fmt.Print("│")
    for x := minX; x <= maxX ; x++ {
      p := &Coord{x, y}
      if matrix[*p] == '#' {
        fmt.Print("\033[48;2;153;153;102m \033[0m")
      } else if matrix[*p] == 'o' {
        fmt.Print("\033[48;2;204;153;0m \033[0m")
      } else if matrix[*p] == '+' {
        fmt.Print("\033[48;2;255;0;0m \033[0m")
      } else {
        fmt.Print("\033[48;2;61;61;41m \033[0m")
      }
    }
    fmt.Println("│") 
  }

  fmt.Print("└")
  for x := minX; x <= maxX ; x++ {
    fmt.Print("─")
  }
  fmt.Println("┘")
}

const FALL = -1 // if a grain of sand fall in the void.
const REST = 0  // if a grain of sand rest in a specific point.
const FULL = 1  // if the source of the sand becomes blocked.

/**
 * Generate a new grain of sand and drop it.
 */
func Step(matrix map[Coord]int8, start *Coord, YLimit int, void bool) int {
  p := &Coord{start.X(), start.Y()}

  for p.Y() < YLimit-1 {

    // can go down
    if _, something := matrix[*p.Down()]; !something {
      p.SetY(p.Y() + 1)
    } else {
    
      if _, something := matrix[*p.DownLeft()]; !something { // try down-left
        p.SetY(p.Y() + 1) 
        p.SetX(p.X() - 1)
      } else if _, something := matrix[*p.DownRight()]; !something { // try bown-right
        p.SetY(p.Y() + 1) 
        p.SetX(p.X() + 1)
      } else {
        matrix[*p] = 'o'
        if *p == *start { // the grain of sand blocked the source
          return FULL
        }
        return REST
      }
    }

  }

  if void { // fall in the void
    return FALL
  } else { // is on the floor
    matrix[*p] = 'o'
    return REST
  }
}

func main() {

  // read input file
  data, err := os.ReadFile("input")
  if err != nil {
    panic(err)
  }
  
  input := strings.Split(string(data), "\n")
  
  // parse the input
  walls := make([][]*Coord, 0)
  for _, row := range input {
    if w, err := ParseRow(row); err == nil {
      walls = append(walls, w)
    }
  }
  
  // find min and max wall's coordinates
  maxX, maxY := 0, 0
  minX, minY := 10000, 10000
  for _, w := range walls {
    for _, c := range w {
      if c.X() > maxX {
        maxX = c.X()
      }
      if c.Y() > maxY {
        maxY = c.Y()
      }

      if c.X() < minX {
        minX = c.X()
      }
      if c.Y() < minY {
        minY = c.Y()
      }
    }
  }
  
  // a map with all wall points
  matrix := GenerateMap(maxY+2, maxX+2, walls)
  PrintMap(matrix, minX-2, 0, maxX+2, maxY+2) // comment here

  x, y := 500, 0 // the coordinate of the source ot the sand

  // compute the result
  result := 0
  for Step(matrix, &Coord{x, y}, maxY+2, true) == REST{
    result += 1
  }

  fmt.Printf("Part One: %d\n", result)
  PrintMap(matrix, minX-2, 0, maxX+2, maxY+2) // comment here

  // --- Part Two ---

  // reset matrix
  matrix = GenerateMap(maxY+2, maxX+2, walls)
  
  // execute the same simulation but with floor
  result = 1
  for Step(matrix, &Coord{x, y}, maxY+2, false) != FULL {
    result += 1
  }

  fmt.Printf("Part Two: %d\n", result)
  
  for k := range matrix {
    if k.X() < minX {
      minX = k.X()
    }
    if k.X() > maxX {
      maxX = k.X()
    }
  }
  PrintMap(matrix, minX-2, 0, maxX+2, maxY+2) // comment here
}
