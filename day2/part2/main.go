package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	h int
	d int
	a int
}

func (p position) move(moveType string, distance int) position {

	switch moveType {
	case "forward":
		return position{
			h: p.h + distance,
			d: p.d + p.a*distance,
			a: p.a,
		}

	case "down":
		return position{
			h: p.h,
			d: p.d,
			a: p.a + distance,
		}

	case "up":
		return position{
			h: p.h,
			d: p.d,
			a: p.a - distance,
		}

	default:
		fmt.Printf("warning invalid moveType: %q\n", moveType)
		return p
	}
}

func part1(fname string) (position, error) {
	f, err := os.Open(fname)
	if err != nil {
		return position{}, err
	}

	defer f.Close()

	p := position{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		v := scanner.Text()
		parts := strings.Split(v, " ")
		if len(parts) != 2 {
			return position{}, fmt.Errorf("invalid input: %q", v)
		}

		d, err := strconv.Atoi(parts[1])
		if err != nil {
			return position{}, fmt.Errorf("failed to parse int value from %q", parts[1])
		}

		p = p.move(parts[0], d)
	}

	return p, nil
}

func main() {
	p, err := part1("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("position: %#v\nresult: %d", p, p.d*p.h)
}
