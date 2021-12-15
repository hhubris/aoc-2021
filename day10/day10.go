package day10

import (
	"bufio"
	"fmt"
	"os"
)

func ReadData(fname string) ([][]rune, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", fname, err)
	}

	defer f.Close()

	var result [][]rune
	s := bufio.NewScanner(f)
	for s.Scan() {
		result = append(result, []rune(s.Text()))
	}

	if err = s.Err(); err != nil {
		return nil, fmt.Errorf("failed to parse entire file: %w", err)
	}

	return result, nil
}

var (
	formatScoreMap = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	missingScoreMap = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

type FormatError struct {
	Expected string
	Found    string
}

func (f *FormatError) Error() string {
	return fmt.Sprintf("expected: %s, found: %s", f.Expected, f.Found)
}

func (f *FormatError) Is(err error) bool {
	if _, ok := err.(*FormatError); ok {
		return true
	}
	return false
}

func (f *FormatError) Score() int {
	return formatScoreMap[f.Found]
}

type MissingMatches struct {
	Missing string
}

func (mm *MissingMatches) Error() string {
	return fmt.Sprintf("missing matches for: %s", mm.Missing)
}

func (mm *MissingMatches) Is(err error) bool {
	if _, ok := err.(*MissingMatches); ok {
		return true
	}
	return false
}

func (mm *MissingMatches) Score() int {
	tot := 0
	for i := len(mm.Missing) - 1; i >= 0; i-- {
		r := rune(mm.Missing[i])
		tot = tot*5 + missingScoreMap[delim[r]]
	}

	return tot
}

var (
	delim = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
)

func validate(i int, r, p []rune) (int, error) {
	if i >= len(r) {
		return 0, &MissingMatches{Missing: string(p)}
	}

	curr := r[i]
	// fmt.Printf("i: %d, c: %s, l: %d, p: %#v\n", i, string(curr), len(r), p)

	if want, ok := delim[curr]; ok {
		ni, err := validate(i+1, r, append(p, curr))
		if err != nil {
			return 0, err
		}

		next := r[ni]
		//fmt.Printf("i: %d, ni: %d, n: %s p: %#v\n", i, ni, string(next), p)
		if next != want {
			return 0, &FormatError{Expected: string(want), Found: string(next)}
		}

		//fmt.Printf("matched %d, %#v\n", next, p)

		if ni == len(r)-1 && len(p) == 0 {
			return 0, nil
		}

		return validate(ni+1, r, p)
	}

	return i, nil
}

func Validate(r []rune) error {
	_, err := validate(0, r, []rune{})
	return err
}
