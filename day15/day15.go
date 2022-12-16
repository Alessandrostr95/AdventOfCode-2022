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

/**
 * Coordinate data structure
 */
type Coord struct {
  X, Y int
}

/**
 * Sensor data structure
 */
type Sensor struct {
  Pos *Coord
  Beacon *Coord
  Distance int
}

/**
 * Copute Manhattan distance
 */
func ManhattanDist(p1, p2 *Coord) int {
  return int(math.Abs(float64(p1.X) - float64(p2.X))) + int(math.Abs(float64(p1.Y) - float64(p2.Y)))
}

/**
 * Parse an input row to a sensor
 */
func ParseRow(s string) (*Sensor, error) {
  if s == "" {
    return nil, errors.New("Invalid Syntax.")
  }

  s = strings.TrimPrefix(s, "Sensor at ")
  s = strings.ReplaceAll(s, " closest beacon is at ", "")
  s = strings.ReplaceAll(s, ", ", ":")
  
  coords := strings.Split(s, ":")

  X1, err := strconv.Atoi(strings.ReplaceAll(coords[0], "x=", ""))
  if err != nil {
    return nil, err
  }

  Y1, err := strconv.Atoi(strings.ReplaceAll(coords[1], "y=", ""))
  if err != nil {
    return nil, err
  }

  X2, err := strconv.Atoi(strings.ReplaceAll(coords[2], "x=", ""))
  if err != nil {
    return nil, err
  }

  Y2, err := strconv.Atoi(strings.ReplaceAll(coords[3], "y=", ""))
  if err != nil {
    return nil, err
  }

  pos := &Coord{X1, Y1}
  beacon := &Coord{X2, Y2}
  dist := ManhattanDist(pos, beacon)

  return &Sensor{pos, beacon, dist}, nil
}

/**
 * Check if `p` is in a range of sensors `s`.
 */
func (s *Sensor) isInRange(p *Coord) bool {
  return ManhattanDist(s.Pos, p) <= s.Distance 
}

/**
 * Return the interval of x-values reachable from sensor `s` at line Y.
 * If `removeBeacon = true` removes the beacon from range.
 */
func (s *Sensor) getInterval(Y int, removeBeacon bool) ([2]int, bool) {
  r := s.Distance - int(math.Abs(float64(s.Pos.Y - Y)))
  if r >= 0 {
    if removeBeacon && s.Beacon.Y == Y && s.Beacon.X <= s.Pos.X {
      return [2]int{s.Pos.X - r + 1, s.Pos.X + r}, true
    } else if removeBeacon && s.Beacon.Y == Y && s.Beacon.X > s.Pos.X {
      return [2]int{s.Pos.X - r, s.Pos.X + r - 1}, true
    }
    return [2]int{s.Pos.X - r, s.Pos.X + r}, true
  } else {
    return [2]int{0, 0}, false // empty interval
  }
}

/**
 * Returns the cardinality of the union of all intervals in input.
 */
func SumIntervals(intervals [][2]int) int {
  if len(intervals) == 0 {
    return 0
  }

  sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
  since := intervals[0][0]
  result := 1
  for i := 0; i < len(intervals); i++ {
    a := intervals[i]
    
    if a[0] > since {
      since = a[0]-1
    }

    if a[1] >= since {
      result += int(math.Abs(float64(a[1] - since)))
      since = a[1]
    }
  }
  return result
}

/**
 * Find x-coord where there is only one space between two intervals.
 */
func Here(intervals [][2]int) int {
  // assuming already sorted
  for i := 0; i < len(intervals) - 1; i++ {
    for j := i + 1; j < len(intervals); j++ {
      a, b := intervals[i], intervals[j]
      if a[1] + 1 == b[0] - 1 {
        return a[1] + 1
      }
    }
  }
  return -1
}

func main() {

  // read input file
  data, err :=  os.ReadFile("input")
  if err != nil {
    panic(err)
  }

  input := strings.Split(string(data), "\n")
  
  // parse all sensors
  sensors := make([]*Sensor, 0)
  for _, row := range input {
    if s, err := ParseRow(row); err == nil {
      sensors = append(sensors, s)
    }
  }

  // --- Part One ---
  
  // find all intervals on line Y
  Y := 2_000_000
  intervals := make([][2]int, 0)
  for _, s := range sensors {
    if i, ok := s.getInterval(Y, true); ok {
      intervals = append(intervals, i)
    }
  }

  result := SumIntervals(intervals) // sum their union
  fmt.Printf("Part One: %d\n", result)

  // --- Part Two ---
  m := 4_000_000
  for y := 0; y <= m; y++ {

    // get intervals at line y
    intervals := make([][2]int, 0)
    for _, s := range sensors {
      if i, ok := s.getInterval(y, false); ok {
        intervals = append(intervals, i)
      }
    }

    // find x-coord where there is only one space between two intervals
    if x := Here(intervals); x != -1 {
      // check if is out of all intervals
      isIsolate := true 
      for _, I := range intervals {
        isIsolate = isIsolate && !(I[0] <= x && x <= I[1])
      }
      
      if isIsolate {
        fmt.Printf("Part Two: %d\n", x*m + y)
        return
      }
    }
  }
 }
