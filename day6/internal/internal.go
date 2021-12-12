package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	spawnTimer    = 6
	newSpawnTimer = 8
)

// i will freely admit, my first attempt at this was array based, as the problem suggested
// part2 quickly showed me that wasn't scalable, and I realized we don't care about individual
// laternfish, just how many exist with the same number of days before they spawn again

func NextGeneration(cg map[int]int) map[int]int {
	result := map[int]int{}
	for k, v := range cg {
		if k == 0 {
			result[newSpawnTimer] += v
			result[spawnTimer] += v
		} else {
			result[k-1] += v
		}
	}

	return result
}

func ReadData(fname string) (map[int]int, error) {
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
	result := map[int]int{}

	for _, v := range parts {
		dv, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("failed to convert %q to an int: %w", v, err)
		}

		result[dv]++
	}

	return result, nil
}

func CountFish(cg map[int]int) int {
	tot := 0

	for _, v := range cg {
		tot += v
	}

	return tot
}
