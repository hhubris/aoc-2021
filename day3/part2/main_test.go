package main

import "testing"

func Test_toDiagnostic(t *testing.T) {

	want := diagnostic{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}
	got := toDiagnostic("101010101010")

	if want != got {
		t.Errorf("want: %#v, got %#v", want, got)
	}
}

func Test_diagnostic_value(t *testing.T) {
	tests := []struct {
		name string
		d    diagnostic
		want int
	}{
		{
			name: "zero",
			d:    diagnostic{},
			want: 0,
		},
		{
			name: "one",
			d:    diagnostic{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			want: 1,
		},
		{
			name: "three",
			d:    diagnostic{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.value()
			if tt.want != got {
				t.Errorf("want: %d, got: %d", tt.want, got)
			}
		})
	}
}

func Test_filter(t *testing.T) {
	diagnostics := []diagnostic{
		toDiagnostic("00100"),
		toDiagnostic("11110"),
		toDiagnostic("10110"),
		toDiagnostic("10111"),
		toDiagnostic("10101"),
		toDiagnostic("01111"),
		toDiagnostic("00111"),
		toDiagnostic("11100"),
		toDiagnostic("10000"),
		toDiagnostic("11001"),
		toDiagnostic("00010"),
		toDiagnostic("01010"),
	}

	tests := []struct {
		name string
		rt   ratingType
		want diagnostic
	}{
		{
			name: "oxy",
			rt:   oxygen,
			want: toDiagnostic("10111"),
		},
		{
			name: "co2",
			rt:   co2,
			want: toDiagnostic("01010"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := filter(diagnostics, tt.rt)
			if err != nil {
				t.Fatal(err)
			}
			if tt.want != got {
				t.Errorf("want: %#v, got: %#v", tt.want, got)
			}
		})
	}
}
