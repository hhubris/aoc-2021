package day11

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type cell struct {
	v       int
	flashed bool
}

func addOne(d [][]cell) {
	for _, row := range d {
		for j, vv := range row {
			row[j].v = vv.v + 1
		}
	}
}

func dump(d [][]cell) {
	for i, v := range d {
		fmt.Printf("%d: %#v\n", i, v)
	}

	fmt.Println()
}

func processFlash(d [][]cell, i, j int) int {
	// fmt.Printf("i: %d j: %d\n", i, j)
	d[i][j].flashed = true

	tot := 1
	for y := i - 1; y <= i+1; y++ {
		for x := j - 1; x <= j+1; x++ {
			if y >= 0 && y < len(d) &&
				x >= 0 && x < len(d[y]) {
				d[y][x].v++

				// dump(d)

				if d[y][x].v > 9 && !d[y][x].flashed {
					tot += processFlash(d, y, x)
				}
			}
		}
	}
	return tot
}

func processFlashes(d [][]cell) {
	for {
		flashed := false
		for i := range d {
			for j := range d[i] {
				if d[i][j].v > 9 && !d[i][j].flashed {
					flashed = true
					processFlash(d, i, j)
				}
			}
		}

		if !flashed {
			break
		}
	}
}

func resetFlashed(d [][]cell) int {
	tot := 0
	for i := range d {
		for j := range d[i] {
			if d[i][j].v > 9 {
				d[i][j].flashed = false
				d[i][j].v = 0
				tot++
			}
		}
	}

	return tot
}

func RunStep(d [][]cell) int {

	addOne(d)
	processFlashes(d)
	return resetFlashed(d)
}

func ParseData(d []byte) ([][]cell, error) {
	var result [][]cell

	s := bufio.NewScanner(bytes.NewReader(d))

	for s.Scan() {
		r := make([]cell, len(s.Text()))
		for i, v := range s.Text() {
			r[i] = cell{v: int(v - '0')}
		}

		result = append(result, r)
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan entire input: %w", err)
	}

	return result, nil
}

func ReadData(fname string) ([][]cell, error) {
	d, err := os.ReadFile(fname)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", fname, err)
	}

	return ParseData(d)
}
