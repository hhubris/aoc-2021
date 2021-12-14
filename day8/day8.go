package day8

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	inputCount  = 10
	outputCount = 4
)

// 0 - 6 signals
// 1 - 2 signals
// 2 - 5 signals
// 3 - 5 signals
// 4 - 4 signals
// 5 - 5 signals
// 6 - 6 signals
// 7 - 3 signals
// 8 - 7 signals
// 9 - 6 signals

type Signal []rune

func (s Signal) includesAll(r []rune) bool {
	sm := make(map[rune]struct{}, len(r))
	for _, x := range s {
		sm[x] = struct{}{}
	}

	for _, x := range r {
		if _, ok := sm[x]; !ok {
			return false
		}
	}

	return true
}

func (s Signal) isZero(r []rune) bool {
	return len(s) == 6 && !s.includesAll(r)
}

func (s Signal) isOne() bool {
	return len(s) == 2
}

func (s Signal) isTwo(r []rune) bool {
	return len(s) == 5 && !s.includesAll(r)
}

func (s Signal) isThree(r []rune) bool {
	return len(s) == 5 && s.includesAll(r)
}

func (s Signal) isFour() bool {
	return len(s) == 4
}

func (s Signal) isFive(r []rune) bool {
	return len(s) == 5 && !s.includesAll(r)
}

func (s Signal) isSeven() bool {
	return len(s) == 3
}

func (s Signal) isEight() bool {
	return len(s) == 7
}

func (s Signal) isNine(r []rune) bool {
	return len(s) == 6 && s.includesAll(r)
}

func (s Signal) isUnique() bool {
	return s.isOne() || s.isFour() || s.isSeven() || s.isEight()
}

type Data struct {
	Input  [inputCount]Signal
	Output [outputCount]Signal
}

func ReadData(fname string) ([]Data, error) {

	f, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", fname, err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	var data []Data
	for s.Scan() {
		parts := strings.Fields(s.Text())
		if len(parts) != 15 {
			return nil, fmt.Errorf("invalid input format: %s", s.Text())
		}

		d := Data{}
		for i := 0; i < inputCount; i++ {
			v := []rune(parts[i])
			sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
			d.Input[i] = v
		}

		for i := 0; i < outputCount; i++ {
			v := []rune(parts[i+inputCount+1])
			sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
			d.Output[i] = v
		}

		data = append(data, d)
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("failed to read input file: %w", err)
	}

	return data, nil
}

func CountUniqueOutputs(d []Data) int {
	tot := 0

	for _, v := range d {
		for _, o := range v.Output {
			if o.isUnique() {
				tot++
			}
		}
	}

	return tot
}

func difference(a, b []rune) []rune {
	mb := make(map[rune]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}

	var diff []rune
	for _, x := range a {
		if _, ok := mb[x]; !ok {
			diff = append(diff, x)
		}
	}

	return diff
}

// i'm not sure how to approach this
// my goal is to return the inputs in a slice such that s[0] is the Signal for 0, s[1] is for 1, etc
func GetOrderedSignals(d Data) [inputCount]Signal {

	result := [inputCount]Signal{}

	// first, let's rearrange the inputs into a hash, so we can remove the ones we've already used easily
	// along the way, save and drop the easy ones
	sm := map[int]Signal{}
	for i, v := range d.Input {
		if v.isOne() {
			result[1] = v
		} else if v.isFour() {
			result[4] = v
		} else if v.isSeven() {
			result[7] = v
		} else if v.isEight() {
			result[8] = v
		} else {
			sm[i] = v
		}
	}

	// next, we can find 3, 9
	for k, v := range sm {
		// 3 has all of 7, and 6 signals
		if v.isThree(result[7]) {
			result[3] = v
			delete(sm, k)
		}

		// 9 has all of 7, and 6 signals
		if v.isNine(result[4]) {
			result[9] = v
			delete(sm, k)
		}
	}

	// start isolating segments using set difference
	top := difference(result[7], result[1])
	bottom := difference(difference(result[9], result[4]), top)
	middle := difference(difference(result[3], result[7]), bottom)
	topLeft := difference(difference(result[4], result[1]), middle)

	// now we can find 0 and 2
	for k, v := range sm {
		// 0 does not have a middle signal, but has all the rest
		if v.isZero(middle) {
			result[0] = v
			delete(sm, k)
		}

		// and of what is left, 2 is the only one that does not have a top left signal
		if v.isTwo(topLeft) {
			result[2] = v
			delete(sm, k)
		}
	}

	// and a final loop to get 5 and 6
	for k, v := range sm {
		if len(v) == 5 {
			result[5] = v
			delete(sm, k)
		}

		if len(v) == 6 {
			result[6] = v
			delete(sm, k)
		}
	}

	if len(sm) != 0 {
		panic(fmt.Sprintf("values left over: %#v", sm))
	}

	return result
}

func GetDisplayValue(input [inputCount]Signal, output [outputCount]Signal) int {
	// first, make a simple map to make lookups easier
	sm := map[string]int{}
	for i, v := range input {
		sm[string(v)] = i
	}

	mult := 1000
	tot := 0

	for _, v := range output {
		tot += sm[string(v)] * mult
		mult = mult / 10
	}

	return tot
}
