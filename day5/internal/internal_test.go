package internal

import (
	"reflect"
	"testing"
)

func TestReadData(t *testing.T) {
	want := []Pair{
		{v1: Coord{x: 0, y: 9}, v2: Coord{x: 5, y: 9}},
		{v1: Coord{x: 8, y: 0}, v2: Coord{x: 0, y: 8}},
		{v1: Coord{x: 9, y: 4}, v2: Coord{x: 3, y: 4}},
		{v1: Coord{x: 2, y: 2}, v2: Coord{x: 2, y: 1}},
		{v1: Coord{x: 7, y: 0}, v2: Coord{x: 7, y: 4}},
		{v1: Coord{x: 6, y: 4}, v2: Coord{x: 2, y: 0}},
		{v1: Coord{x: 0, y: 9}, v2: Coord{x: 2, y: 9}},
		{v1: Coord{x: 3, y: 4}, v2: Coord{x: 1, y: 4}},
		{v1: Coord{x: 0, y: 0}, v2: Coord{x: 8, y: 8}},
		{v1: Coord{x: 5, y: 5}, v2: Coord{x: 8, y: 2}},
	}

	got, err := ReadData("../test_data.txt")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %#v\n got: %#v\n", want, got)
	}
}

func TestCreateDataMap(t *testing.T) {

	pairs, err := ReadData("../test_data.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := CreateDataMap(pairs, false)
	ints := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 1, 2, 1, 1, 1, 2, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
	}

	want := DataMap{}

	for y, row := range ints {
		for x, v := range row {
			want[y][x] = v
		}
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %#v\n got:%#v\n", want, got)
	}
}

func Test_DataMap_CountDangerous(t *testing.T) {

	pairs, err := ReadData("../test_data.txt")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name        string
		includeDiag bool
		want        int
	}{
		{
			name:        "no diag",
			includeDiag: false,
			want:        5,
		},
		{
			name:        "yes diag",
			includeDiag: true,
			want:        12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := CreateDataMap(pairs, tt.includeDiag)
			got := dm.CountDangerous(1)

			if tt.want != got {
				t.Errorf("want: %d, got: %d", tt.want, got)
			}

			for i := 0; i < 10; i++ {
				t.Logf("%#v\n", dm[i][0:10])
			}

		})
	}

}
