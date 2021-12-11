package main

import (
	"bufio"
	"fmt"
	"os"
)

type diagnostic [12]int

func (d diagnostic) value() int {

	tot := 0

	for i, v := range d {
		tot += v * powers[i]
	}

	return tot
}

type report struct {
	totals diagnostic
	cnt    int
}

var powers = diagnostic{2048, 1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}

func (r report) add(d diagnostic) report {

	nd := diagnostic{}

	for i, v := range d {
		nd[i] = r.totals[i] + v
	}

	return report{
		totals: nd,
		cnt:    r.cnt + 1,
	}

}

func (r report) gamma() diagnostic {
	d := diagnostic{}

	cutoff := r.cnt / 2
	for i, v := range r.totals {
		val := 0
		if v > cutoff {
			val = 1
		}

		d[i] = val
	}

	return d
}

func (r report) epsilon() diagnostic {
	g := r.gamma()
	d := diagnostic{}

	for i, v := range g {
		if v == 0 {
			d[i] = 1
		}
	}

	return d
}

func toDiagnostic(s string) diagnostic {

	d := diagnostic{}

	for i, v := range s {
		d[i] = int(v - '0')

		// in production code we would check the value is in the expected range
	}

	return d
}

func part1(fname string) (report, error) {
	f, err := os.Open(fname)
	if err != nil {
		return report{}, err
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	r := report{}

	for s.Scan() {
		d := toDiagnostic(s.Text())
		r = r.add(d)
	}
	if err := s.Err(); err != nil {
		return report{}, err
	}

	return r, nil
}

func main() {
	r, err := part1("../input.txt")
	if err != nil {
		panic(err)
	}

	g := r.gamma()
	e := r.epsilon()

	fmt.Printf(" report: %#v\n", r)
	fmt.Printf("  gamma: %#v value: %d\n", g, g.value())
	fmt.Printf("epsilon: %#v value: %d\n", e, e.value())
	fmt.Printf("power consumtion: %d\n", g.value()*e.value())
}
