package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	inputMatcher = regexp.MustCompile(`^(?P<x1>\d+),(?P<y1>\d+)\s+->\s+(?P<x2>\d+),(?P<y2>\d+)$`)
)

type DataMap [1000][1000]int

func (dm DataMap) CountDangerous(min int) int {
	tot := 0
	for _, y := range dm {
		for _, v := range y {
			if v > min {
				tot++
			}
		}
	}

	return tot
}

type Coord struct {
	x int
	y int
}

func (c Coord) String() string {
	return fmt.Sprintf("x: %d, y: %d", c.x, c.y)
}

type Pair struct {
	v1 Coord
	v2 Coord
}

func (p Pair) String() string {
	return fmt.Sprintf("%s -> %s", p.v1, p.v2)
}

func (p Pair) isHorizontal() bool {
	return p.v1.y == p.v2.y
}

func (p Pair) isVertical() bool {
	return p.v1.x == p.v2.x
}

func parseInt(s string) int {
	// normally I wouldn't do this, but since I used a regexp to match digits we shouldn't have any conversion errors
	v, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("WARNING: Failed to parse %q to an int", s)
		return 0
	}

	return v
}

func parseLine(s string) (Pair, error) {
	sm := inputMatcher.FindStringSubmatch(s)
	if len(sm) != 5 {
		return Pair{}, fmt.Errorf("did not find the expected values in : %s", s)
	}

	return Pair{
		v1: Coord{
			x: parseInt(sm[1]),
			y: parseInt(sm[2]),
		},
		v2: Coord{
			x: parseInt(sm[3]),
			y: parseInt(sm[4]),
		},
	}, nil
}

func getOrdered(a, b int) (int, int) {
	if a > b {
		return b, a
	}

	return a, b
}
func addHorizontalLine(dm DataMap, p Pair) DataMap {

	x1, x2 := getOrdered(p.v1.x, p.v2.x)
	for x := x1; x <= x2; x++ {
		dm[p.v1.y][x]++
	}

	return dm
}

func addVerticalLine(dm DataMap, p Pair) DataMap {

	y1, y2 := getOrdered(p.v1.y, p.v2.y)
	for y := y1; y <= y2; y++ {
		dm[y][p.v1.x]++
	}

	return dm
}

func intAbs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func getVector(a, b int) (int, int) {
	dir := 1
	if a > b {
		dir = -1
	}

	return a, dir
}

func addDiagonalLine(dm DataMap, p Pair) DataMap {

	// first figure out how many points we are going to update
	points := intAbs(p.v1.x-p.v2.x) + 1

	x, xofs := getVector(p.v1.x, p.v2.x)
	y, yofs := getVector(p.v1.y, p.v2.y)

	for i := 0; i < points; i++ {
		dm[y][x]++
		x += xofs
		y += yofs
	}

	return dm
}

func addLine(dm DataMap, p Pair, includeDiag bool) DataMap {
	if p.isHorizontal() {
		return addHorizontalLine(dm, p)
	} else if p.isVertical() {
		return addVerticalLine(dm, p)
	} else if includeDiag {
		return addDiagonalLine(dm, p)
	}

	return dm
}

func CreateDataMap(p []Pair, includeDiag bool) DataMap {

	dm := DataMap{}

	for _, v := range p {
		dm = addLine(dm, v, includeDiag)
	}

	return dm
}

func ReadData(fname string) ([]Pair, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	var p []Pair
	for s.Scan() {
		np, err := parseLine(s.Text())
		if err != nil {
			return nil, fmt.Errorf("failed to parse input: %s : %w", s.Text(), err)
		}

		p = append(p, np)
	}

	if err = s.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %w", err)
	}

	return p, nil
}
