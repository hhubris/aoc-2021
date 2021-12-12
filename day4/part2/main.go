package main

import (
	"fmt"
	"github.com/hhubris/aoc-2021/day4/internal"
)

func callBingo(c internal.Calls, boards []*internal.Board) (*internal.Board, int) {

	var lastWinner *internal.Board
	var lastCall int

	for _, call := range c {
		for _, b := range boards {
			if !b.HasBingo() && internal.MarkBoard(b, call) {
				b.SetBingo(true)
				lastWinner = b
				lastCall = call
			}
		}
	}

	return lastWinner, lastCall
}

func main() {
	win, lastCall, err := internal.RunBingo("../input.txt", callBingo)
	if err != nil {
		panic(err)
	}

	su := internal.SumUnmarked(win)
	fmt.Printf("lastCall: %d\n", lastCall)
	fmt.Printf("total unmarked: %d\n", su)
	fmt.Printf("final score: %d\n", lastCall*su)
}
