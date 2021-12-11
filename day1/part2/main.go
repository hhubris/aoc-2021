package main

import (
	"bufio"
	"strconv"
	"fmt"
	"os"
)

type triple [3]int

func (t triple) new(a int) triple {
	return triple{a, t[0], t[1]}
}

func scanFirstTriple(s *bufio.Scanner) (triple, error) {
	vals := triple{}
	for i := 0; i < 3; i++ {
		if !s.Scan() {
			return triple{}, fmt.Errorf("failed to find enough values in input")
		}
		
		if err := s.Err(); err != nil {
			return triple{}, err
		}

		v, err := strconv.Atoi(s.Text())
		if err != nil {
			return triple{}, err
		}

		vals[i] = v
	}

	return vals, nil
}

func part2(fname string) (int, error) {
	f, err := os.Open(fname)
	if err != nil {
		return 0, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	
	prev, err := scanFirstTriple(scanner)
	if err != nil {
		return 0, err
	}
	
	cnt := 0
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		
		cur := prev.new(v)
		// by definition cur[1] = prev[0] and cur[2] prev[1], so we only need to compare cur[0] against prev[2=, no need for sums
		if cur[0] > prev[2] {
			cnt++
		}
		
		prev = cur
	}
	
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return cnt, nil
}

func main() {
	cnt, err := part2("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %d increasing triples", cnt)
}
