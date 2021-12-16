package day12

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestFindPaths(t *testing.T) {

	input := []byte(`start-A
start-b
A-c
A-b
b-d
A-end
b-end`)

	want := [][]string{
		{"start", "A", "b", "A", "c", "A", "end"},
		{"start", "A", "b", "A", "end"},
		{"start", "A", "b", "end"},
		{"start", "A", "c", "A", "b", "A", "end"},
		{"start", "A", "c", "A", "b", "end"},
		{"start", "A", "c", "A", "end"},
		{"start", "A", "end"},
		{"start", "b", "A", "c", "A", "end"},
		{"start", "b", "A", "end"},
		{"start", "b", "end"},
	}

	d, err := MakeMap(bytes.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	paths := FindPaths(d)

	t.Logf("%#v\n", paths)

	if !cmp.Equal(want, paths) {
		t.Errorf(cmp.Diff(paths, want))
	}
}

func TestFindPaths2(t *testing.T) {
	input := []byte(`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`)

	d, err := MakeMap(bytes.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	paths := FindPaths(d)

	t.Logf("len: %d", len(paths))

	for i, p := range paths {
		t.Logf("%d: %s", i, strings.Join(p, ","))
	}
}

func TestPart1(t *testing.T) {

	d, err := ReadData("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	paths := FindPaths(d)
	wantCnt := 4167
	gotCnt := len(paths)

	if wantCnt != gotCnt {
		t.Errorf("want: %d, got: %d", wantCnt, gotCnt)
	}

}
