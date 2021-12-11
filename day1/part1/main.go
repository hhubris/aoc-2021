package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(fname string) (int, error) {
	f, err := os.Open(fname)
	if err != nil {
		return 0, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	prev := 0
	cnt := 0
	skip := true

	for scanner.Scan() {
		line := scanner.Text()
		v, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}

		if skip {
			skip = false
			prev = v
			continue
		}

		if v > prev {
			cnt++
		}

		prev = v
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return cnt, nil

}

func main() {
	cnt, err := part1("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %d increasing steps", cnt)
}
