package main

import "testing"

func TestPart1(t *testing.T) {

	cnt, err := part1("testdata/test.txt")
	if err != nil {
		t.Fatal(err)
	}

	if cnt != 7 {
		t.Errorf("want: 7, got %d", cnt)
	}
}
