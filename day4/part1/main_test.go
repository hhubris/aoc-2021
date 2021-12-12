package main

import (
	"github.com/hhubris/aoc-2021/day4/internal"
	"reflect"
	"testing"
)

func Test_callBingo(t *testing.T) {

	c := internal.Calls{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24}

	b := [5][5]int{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	}
	want := internal.MakeBoard(b)

	got, lastCall, err := internal.RunBingo("../test_data.txt", callBingo)
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range c {
		internal.MarkBoard(want, c)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %#v\n got: %#v\n", want, got)
	}

	finalVal := lastCall * internal.SumUnmarked(got)

	if 4512 != finalVal {
		t.Errorf("want: 4512, got: %d", finalVal)
	}
}
