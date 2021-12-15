package day9

import (
	"reflect"
	"testing"
)

func TestReadData(t *testing.T) {
	want := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	got, err := ReadData("test_data.txt")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestFindMinimumDepths(t *testing.T) {
	data, err := ReadData("test_data.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := []Point{
		{X: 1, Y: 0, H: 1},
		{X: 9, Y: 0, H: 0},
		{X: 2, Y: 2, H: 5},
		{X: 6, Y: 4, H: 5},
	}

	got := FindMinimumDepths(data)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nwant: %#v\n got: %#v", want, got)
	}
}

func TestFindTotalRiskLevel(t *testing.T) {
	tests := []struct {
		name string
		fn   string
		want int
	}{
		{
			name: "test",
			fn:   "test_data.txt",
			want: 15,
		},
		{
			name: "data",
			fn:   "input.txt",
			want: 462,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ReadData(tt.fn)
			if err != nil {
				t.Fatal(err)
			}

			mins := FindMinimumDepths(data)

			got := TotalRiskLevel(mins)
			if tt.want != got {
				t.Errorf("want: %d, got: %d", tt.want, got)
			}

		})
	}
}

func TestFindBasinSize(t *testing.T) {
	data, err := ReadData("test_data.txt")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		min  Point
		want int
	}{
		{
			name: "1x0",
			min:  Point{X: 1, Y: 0, H: 1},
			want: 3,
		}, {
			name: "9x0",
			min:  Point{X: 9, Y: 0, H: 0},
			want: 9,
		}, {
			name: "2x2",
			min:  Point{X: 2, Y: 2, H: 5},
			want: 14,
		}, {
			name: "6x4",
			min:  Point{X: 6, Y: 4, H: 5},
			want: 9,
		},
	}

	for _, tt := range tests {
		got := FindBasinSize(data, tt.min)
		if tt.want != got {
			t.Errorf("want: %d, got: %d", tt.want, got)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name string
		fn   string
		want int
	}{
		{
			name: "test",
			fn:   "test_data.txt",
			want: 1134,
		},
		{
			name: "part2",
			fn:   "input.txt",
			want: 1397760,
		},
	}

	for _, tt := range tests {
		data, err := ReadData(tt.fn)
		if err != nil {
			t.Fatal(err)
		}

		got := Part2(data)
		if tt.want != got {
			t.Errorf("want: %d, got: %d", tt.want, got)
		}
	}
}
