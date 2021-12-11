package main

import "testing"

func Test_toDiagnostic(t *testing.T) {

	want := diagnostic{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}
	got := toDiagnostic("101010101010")

	if want != got {
		t.Errorf("want: %#v, got %#v", want, got)
	}
}

func Test_report_add(t *testing.T) {

	d := []diagnostic{
		{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	}

	want := []report{
		{
			totals: diagnostic{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			cnt:    1,
		},
		{
			totals: diagnostic{1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			cnt:    2,
		},
	}

	r := report{}
	for i, v := range d {
		r = r.add(v)

		if r != want[i] {
			t.Errorf("want: %#v\n, got: %#v", want[i], r)
		}
	}
}

func Test_report_gamma(t *testing.T) {

	tests := []struct {
		name string
		r    report
		want diagnostic
	}{
		{
			name: "empty",
			r:    report{},
			want: diagnostic{},
		},
		{
			name: "first only",
			r: report{
				totals: diagnostic{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				cnt:    1,
			},
			want: diagnostic{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "some",
			r: report{
				totals: diagnostic{3, 1, 2, 1, 0, 4, 1, 2, 3, 5, 0, 0},
				cnt:    5,
			},
			want: diagnostic{1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.gamma()
			if tt.want != got {
				t.Errorf("want: %#v, got: %#v", tt.want, got)
			}
		})
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
