package day9

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func parseRow(s string) ([]int, error) {

	result := make([]int, len(s))
	for i, v := range s {
		result[i] = int(v - '0')
	}

	return result, nil
}

func ReadData(fname string) ([][]int, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", fname, err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	var result [][]int
	for s.Scan() {
		r, err := parseRow(s.Text())
		if err != nil {
			return nil, err
		}

		result = append(result, r)
	}

	if err = s.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan entire file: %w", err)
	}

	return result, nil
}

type Point struct {
	X int
	Y int
	H int
}

func isMinimum(d [][]int, y, x int) bool {
	v := d[y][x]

	if (y > 0 && v >= d[y-1][x]) || // look up
		(y < len(d)-1 && v >= d[y+1][x]) || //look down
		(x > 0 && v >= d[y][x-1]) || // look left
		(x < len(d[y])-1 && v >= d[y][x+1]) { // look right
		return false
	}

	return true
}

func FindMinimumDepths(d [][]int) []Point {

	var result []Point

	for y := 0; y < len(d); y++ {
		for x := 0; x < len(d[y]); x++ {
			if isMinimum(d, y, x) {
				result = append(result, Point{X: x, Y: y, H: d[y][x]})
			}
		}
	}

	return result
}

func TotalRiskLevel(p []Point) int {
	tot := 0

	for _, v := range p {
		tot += 1 + v.H
	}

	return tot
}

type PointMap map[Point]struct{}

func findBasinSize(d [][]int, min Point, basin PointMap) int {

	// if the current point is already in the map, or the height of the point is 9, we don't include the point
	if _, ok := basin[min]; ok || min.H == 9 {
		return 0
	}

	// start at 1 since we know the min point has to count as part of the size of the basin
	tot := 1
	basin[min] = struct{}{}

	// now add all the points that can flow into the current point

	// flow right
	if min.X > 0 && d[min.Y][min.X-1] > min.H {
		tot += findBasinSize(d, Point{X: min.X - 1, Y: min.Y, H: d[min.Y][min.X-1]}, basin)
	}

	// flow left
	if min.X < len(d[min.Y])-1 && d[min.Y][min.X+1] > min.H {
		tot += findBasinSize(d, Point{X: min.X + 1, Y: min.Y, H: d[min.Y][min.X+1]}, basin)
	}

	// flow down
	if min.Y > 0 && d[min.Y-1][min.X] > min.H {
		tot += findBasinSize(d, Point{X: min.X, Y: min.Y - 1, H: d[min.Y-1][min.X]}, basin)
	}

	// flow up
	if min.Y < len(d)-1 && d[min.Y+1][min.X] > min.H {
		tot += findBasinSize(d, Point{X: min.X, Y: min.Y + 1, H: d[min.Y+1][min.X]}, basin)
	}

	return tot
}

func FindBasinSize(d [][]int, min Point) int {
	// again, I'm not sure how to approach this
	// I think I need to start at the min and figure out how wide and tall the bounding box of the basin is
	// then figure out all the points that can flow into those

	// oh, different idea, what if i just check the 4 points around the target to find out if they can flow to the target
	// then repeat, using all points that can flow to the original target, as the new target
	// my brain is telling me recursion is going to be helpful here

	return findBasinSize(d, min, PointMap{})
}

func Part2(d [][]int) int {
	mins := FindMinimumDepths(d)

	sizes := make([]int, len(mins))

	for i, v := range mins {
		sizes[i] = FindBasinSize(d, v)
	}

	sort.Slice(sizes, func(a, b int) bool { return sizes[b] < sizes[a] })
	return sizes[0] * sizes[1] * sizes[2]

}
