package main

import (
	"fmt"
	"github.com/hhubris/aoc-2021/day6/internal"
)

func main() {

	cg, err := internal.ReadData("../input.txt")
	if err != nil {
		panic(err)
	}

	const days = 80
	for i := 0; i < days; i++ {
		cg = internal.NextGeneration(cg)
	}

	fmt.Printf("After %d days: %d lanternfish\n", days, internal.CountFish(cg))

}
