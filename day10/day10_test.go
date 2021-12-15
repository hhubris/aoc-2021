package day10

import (
	"errors"
	"fmt"
	"sort"
	"testing"
)

func TestValidate(t *testing.T) {

	tests := []struct {
		name    string
		v       string
		wantErr bool
	}{
		{
			name:    "v1",
			v:       "([])",
			wantErr: false,
		},
		{
			name:    "v2",
			v:       "{()()()}",
			wantErr: false,
		},
		{
			name:    "v3",
			v:       "<([{}])>",
			wantErr: false,
		},
		{
			name:    "v4",
			v:       "[<>({}){}[([])<>]]",
			wantErr: false,
		},
		{
			name:    "v5",
			v:       "(((((((((())))))))))",
			wantErr: false,
		},
		{
			name:    "v6",
			v:       "[({(<(())[]>[[{[]{<()<>>",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("\nstarting test: %q\n", tt.v)
			gotErr := Validate([]rune(tt.v))

			if tt.wantErr != (gotErr != nil) {
				t.Errorf("wantErr: %v, gotErr: %v", tt.wantErr, gotErr)
			}
		})
	}
}

func TestValidate2(t *testing.T) {

	wantErr := []error{
		&MissingMatches{Missing: "[({([[{{"},
		&MissingMatches{Missing: "({[<{("},
		&FormatError{Expected: "]", Found: "}"},
		&MissingMatches{Missing: "((((<{<{{"},
		&FormatError{Expected: "]", Found: ")"},
		&FormatError{Expected: ")", Found: "]"},
		&MissingMatches{Missing: "<{[{[{{[["},
		&FormatError{Expected: ">", Found: ")"},
		&FormatError{Expected: "]", Found: ">"},
		&MissingMatches{Missing: "<{(["},
	}

	data, err := ReadData("test_data.txt")
	if err != nil {
		t.Fatal(err)
	}

	for i, v := range data {
		t.Run(fmt.Sprintf("v%d", i+1), func(t *testing.T) {
			t.Logf("\nstarting test: %q\n", string(v))

			gotErr := Validate(v)

			if !errors.Is(wantErr[i], gotErr) {
				t.Errorf("i: %d, wantErr: %v, gotErr: %v", i, wantErr[i], gotErr)
			}
		})
	}
}

func TestFormatErrorScore(t *testing.T) {

	tests := []struct {
		name string
		fn   string
		want int
	}{
		{
			name: "test",
			fn:   "test_data.txt",
			want: 26397,
		},
		{
			name: "real",
			fn:   "input.txt",
			want: 278475,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ReadData(tt.fn)
			if err != nil {
				t.Fatal(err)
			}

			got := 0
			var e *FormatError

			for _, v := range data {
				err = Validate(v)
				if errors.As(err, &e) {
					got += e.Score()
				}

			}

			if tt.want != got {
				t.Errorf("want: %d, got %d", tt.want, got)
			}

		})
	}
}

func TestMissingScore(t *testing.T) {

	tests := []struct {
		name string
		fn   string
		want int
	}{
		{
			name: "test",
			fn:   "test_data.txt",
			want: 288957,
		},
		{
			name: "real",
			fn:   "input.txt",
			want: 3015539998,
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ReadData(tt.fn)
			if err != nil {
				t.Fatal(err)
			}

			var e *MissingMatches
			var scores []int
			for _, v := range data {
				err = Validate(v)
				if errors.As(err, &e) {
					scores = append(scores, e.Score())
				}
			}

			sort.Ints(scores)
			got := scores[len(scores)/2]

			if tt.want != got {
				t.Errorf("want: %d, got %d", tt.want, got)
			}

		})
	}
}
