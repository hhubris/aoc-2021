package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cell struct {
	val    int
	marked bool
}
type Board struct {
	cells [5][5]cell
	bingo bool // board already has a bingo
}

func (b *Board) SetBingo(nb bool) {
	b.bingo = nb
}

func (b *Board) HasBingo() bool {
	return b.bingo
}

type Calls []int

func getCalls(s string) (Calls, error) {
	parts := strings.Split(s, ",")

	c := make(Calls, len(parts))

	for i, v := range parts {
		dv, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		c[i] = dv
	}

	return c, nil
}

func getBoard(s *bufio.Scanner) (*Board, error) {

	b := Board{}
	for i := 0; i < 6; i++ {
		if !s.Scan() {
			if i == 0 {
				return nil, nil
			}

			return nil, fmt.Errorf("scan did not return an expected row")
		}

		if i == 0 {
			continue
		}

		parts := strings.Fields(s.Text())
		if len(parts) != 5 {
			return nil, fmt.Errorf("unexpected number of values (%d) found in %s", len(parts), s.Text())
		}

		var row [5]cell
		for i, v := range parts {
			dv, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("unable to parse %q to an int", v)
			}

			row[i] = cell{val: dv}
		}

		b.cells[i-1] = row
	}

	return &b, nil

}

func readData(fname string) (Calls, []*Board, error) {

	f, err := os.Open(fname)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open input file: %w", err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	var c Calls
	if s.Scan() {
		c, err = getCalls(s.Text())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse calls: %w", err)
		}
	}

	var boards []*Board
	for {
		b, err := getBoard(s)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get board: %w", err)
		}

		if b == nil {
			break
		}

		boards = append(boards, b)
	}

	if err := s.Err(); err != nil {
		return nil, nil, fmt.Errorf("failed to read input file: %w", err)
	}

	return c, boards, nil

}

func checkAcross(b *Board, i int) bool {
	for idx := 0; idx < 5; idx++ {
		if !b.cells[i][idx].marked {
			return false
		}
	}

	return true
}

func checkDown(b *Board, j int) bool {
	for idx := 0; idx < 5; idx++ {
		if !b.cells[idx][j].marked {
			return false
		}
	}

	return true
}

func checkBingo(b *Board, i, j int) bool {
	return checkAcross(b, i) || checkDown(b, j)
}

func MarkBoard(b *Board, call int) bool {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.cells[i][j].val == call {
				b.cells[i][j].marked = true
				return checkBingo(b, i, j)
			}
		}
	}

	return false
}

func SumUnmarked(b *Board) int {
	total := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.cells[i][j].marked {
				total += b.cells[i][j].val
			}
		}
	}

	return total
}

type BingoCaller func(Calls, []*Board) (*Board, int)

func RunBingo(fname string, bc BingoCaller) (*Board, int, error) {
	c, b, err := readData(fname)
	if err != nil {
		return nil, 0, err
	}

	win, lastCall := bc(c, b)
	return win, lastCall, nil
}

func MakeBoard(v [5][5]int) *Board {

	b := Board{}
	for ir, row := range v {
		var br [5]cell
		for iv, val := range row {
			br[iv] = cell{val: val}
		}

		b.cells[ir] = br
	}

	return &b
}
