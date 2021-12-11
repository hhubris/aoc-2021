package main

import (
	"bufio"
	"fmt"
	"os"
)

type ratingType int

const (
	oxygen ratingType = iota
	co2

	diagnosticCnt = 12
)

type diagnostic [diagnosticCnt]int

func (d diagnostic) value() int {

	tot := 0

	for i, v := range d {
		tot += v * powers[i]
	}

	return tot
}

var powers = diagnostic{2048, 1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}

func toDiagnostic(s string) diagnostic {

	d := diagnostic{}

	for i, v := range s {
		d[i] = int(v - '0')

		// in production code we would check the value is in the expected range
	}

	return d
}

func targetValue(total, cnt int, rt ratingType) (int, error) {
	switch rt {
	case oxygen:
		if float64(total) >= float64(cnt)/2.0 {
			return 1, nil
		}

		return 0, nil

	case co2:
		if float64(total) >= float64(cnt)/2.0 {
			return 0, nil
		}

		return 1, nil
	default:
		return 0, fmt.Errorf("invalid rating type: %d", rt)
	}
}

func createTotal(diagnostics []diagnostic, idx int) int {
	total := 0
	for _, d := range diagnostics {
		total += d[idx]

	}

	return total
}

func filter(diagnostics []diagnostic, rt ratingType) (diagnostic, error) {

	curr := diagnostics
	for idx := 0; idx < diagnosticCnt; idx++ {

		total := createTotal(curr, idx)
		want, err := targetValue(total, len(curr), rt)
		if err != nil {
			return diagnostic{}, err
		}

		var keep []diagnostic
		for _, d := range curr {
			if d[idx] == want {
				keep = append(keep, d)
			}
		}

		if len(keep) == 1 {
			return keep[0], nil
		}

		curr = keep
	}

	// must have the same value multiple times
	return curr[0], nil
}

func part2(fname string) (diagnostic, diagnostic, error) {
	f, err := os.Open(fname)
	if err != nil {
		return diagnostic{}, diagnostic{}, err
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	var diagnostics []diagnostic
	for s.Scan() {
		diagnostics = append(diagnostics, toDiagnostic(s.Text()))
	}
	if err := s.Err(); err != nil {
		return diagnostic{}, diagnostic{}, err
	}

	o, err := filter(diagnostics, oxygen)
	if err != nil {
		return diagnostic{}, diagnostic{}, fmt.Errorf("failed to find oxygen rating: %w", err)
	}

	c, err := filter(diagnostics, co2)
	if err != nil {
		return diagnostic{}, diagnostic{}, fmt.Errorf("failed to find co2 rating: %w", err)
	}

	return o, c, nil
}

func main() {
	o, c, err := part2("../input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("oxygen: %#v value: %d\n", o, o.value())
	fmt.Printf("co2: %#v value: %d\n", c, c.value())
	fmt.Printf("life support: %d\n", o.value()*c.value())
}
