package main

import (
	"github.com/hhubris/aoc-2021/day4/internal"
	"reflect"
	"testing"
)

func Test_callBingo(t *testing.T) {

	c := internal.Calls{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	b := [5][5]int{
		{3, 15, 0, 2, 22},
		{9, 18, 13, 17, 5},
		{19, 8, 7, 25, 23},
		{20, 11, 10, 24, 4},
		{14, 21, 16, 12, 6},
	}
	want := internal.MakeBoard(b)

	got, lastCall, err := internal.RunBingo("../test_data.txt", callBingo)
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range c {
		if !want.HasBingo() && internal.MarkBoard(want, c) {
			want.SetBingo(true)
		}
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %#v\n got: %#v\n", want, got)
	}

	finalVal := lastCall * internal.SumUnmarked(got)

	if 1924 != finalVal {
		t.Errorf("want: 1924, got: %d", finalVal)
	}
}
