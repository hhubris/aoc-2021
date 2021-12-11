package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_scanFirstTriple(t *testing.T) {

	f, err := os.Open("testdata/test.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	got, err := scanFirstTriple(s)
	if err != nil {
		t.Fatal(err)
	}

	want := triple{199, 200, 208}

	if want != got {
		t.Errorf("want: %#v, got %#v", want, got)
	}
}

func Test_triple_new(t *testing.T) {
	tests := []struct {
		name string
		t    triple
		want triple
	}{
		{
			name: "empty",
			t:    triple{},
			want: triple{43,0,0},
		},
		{
			name: "123",
			t:    triple{1,2,3},
			want: triple{43,1,2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.new(43)
			if tt.want != got {
				t.Errorf("want: %d, got: %d", tt.want, got)
			}
		})
	}
}
