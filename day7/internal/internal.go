package internal

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func SimpleCost(d []int, target int) int {

	tot := 0
	for k, v := range d {
		tot += absInt(k-target) * v
	}

	return tot
}

var (
	geometricCosts []int
	o              = &sync.Once{}
)

func GeometricCost(d []int, target int) int {
	o.Do(func() {
		geometricCosts = make([]int, len(d))
		c := 0
		for i := 0; i < len(d); i++ {
			geometricCosts[i] = c
			c += i + 1
		}
	})

	tot := 0
	for k, v := range d {
		tot += geometricCosts[absInt(k-target)] * v
	}

	return tot
}

type CostMethod func([]int, int) int

func FindMinimumCost(d []int, cm CostMethod) (int, int) {
	currMin := math.MaxInt
	target := 0

	for i := 0; i < len(d); i++ {
		cost := cm(d, i)
		if cost < currMin {
			currMin = cost
			target = i
		}
	}

	return currMin, target
}

// this is the same ReadData function as the one for day6 - not going to bother unit testing it again
func ReadData(fname string) ([]int, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", fname, err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read input data: %w", err)
	}

	parts := strings.Split(string(b), ",")
	rm := map[int]int{}

	max := 0
	for _, v := range parts {
		dv, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("failed to convert %q to an int: %w", v, err)
		}

		if dv > max {
			max = dv
		}

		rm[dv]++
	}

	result := make([]int, max+1)
	for k, v := range rm {
		result[k] = v
	}

	return result, nil
}
