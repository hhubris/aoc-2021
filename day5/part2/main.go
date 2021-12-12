package main

import (
	"fmt"
	"github.com/hhubris/aoc-2021/day5/internal"
)

func main() {

	pairs, err := internal.ReadData("../input.txt")
	if err != nil {
		panic(err)
	}

	dm := internal.CreateDataMap(pairs, true)
	fmt.Printf("Dangerous spots: %d\n", dm.CountDangerous(1))
}
