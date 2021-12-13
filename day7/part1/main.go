package main

import (
	"fmt"
	"github.com/hhubris/aoc-2021/day7/internal"
)

func main() {

	d, err := internal.ReadData("../input.txt")
	if err != nil {
		panic(err)
	}

	min, target := internal.FindMinimumCost(d, internal.SimpleCost)
	fmt.Printf("Minimum cost is %d to move to %d\n", min, target)
}
