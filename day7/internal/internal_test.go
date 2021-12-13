package internal

import (
	"testing"
)

func TestSimpleCost(t *testing.T) {
	d, err := ReadData("../test_data.txt")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		target int
		want   int
	}{
		{
			name:   "two",
			target: 2,
			want:   37,
		},
		{
			name:   "one",
			target: 1,
			want:   41,
		},
		{
			name:   "three",
			target: 3,
			want:   39,
		},
		{
			name:   "ten",
			target: 10,
			want:   71,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimpleCost(d, tt.target)
			if tt.want != got {
				t.Errorf("want: %d, got: %d", tt.want, got)
			}
		})
	}
}

func TestFindMinimumCost(t *testing.T) {
	d, err := ReadData("../test_data.txt")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name       string
		m          CostMethod
		wantCost   int
		wantTarget int
	}{
		{
			name:       "simple",
			m:          SimpleCost,
			wantCost:   37,
			wantTarget: 2,
		},
		{
			name:       "geometric",
			m:          GeometricCost,
			wantCost:   168,
			wantTarget: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCost, gotTarget := FindMinimumCost(d, tt.m)

			if tt.wantCost != gotCost {
				t.Errorf("want: %d, got %d", tt.wantCost, gotCost)
			}

			if tt.wantTarget != gotTarget {
				t.Errorf("want: %d, got %d", tt.wantTarget, gotTarget)
			}
		})
	}

}
