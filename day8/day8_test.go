package day8

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestCountUniqueOutputs(t *testing.T) {
	tests := []struct {
		fn   string
		want int
	}{
		{
			fn:   "test_data.txt",
			want: 26,
		},
		{
			fn:   "input.txt",
			want: 452,
		},
	}

	for _, tt := range tests {
		t.Run(tt.fn, func(t *testing.T) {
			d, err := ReadData(tt.fn)
			if err != nil {
				t.Fatal(err)
			}

			got := CountUniqueOutputs(d)
			if tt.want != got {
				t.Errorf("want: %d, got: %d", tt.want, got)
			}
		})
	}
}

func TestGetOrderedSignals(t *testing.T) {
	data, err := ReadData("test_data.txt")
	if err != nil {
		t.Fatal(err)
	}

	d := data[0]

	want := [inputCount]Signal{
		[]rune("abdefg"),  //0
		[]rune("be"),      //1
		[]rune("abcdf"),   //2
		[]rune("bcdef"),   //3
		[]rune("bceg"),    //4
		[]rune("cdefg"),   //5
		[]rune("acdefg"),  //6
		[]rune("bde"),     //7
		[]rune("abcdefg"), //8
		[]rune("bcdefg"),  //9
	}

	got := GetOrderedSignals(d)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %s\n got: %s\n", dumpSignals(want), dumpSignals(got))
	}

	wantDisplayValue := 8394
	gotDisplayValue := GetDisplayValue(got, d.Output)
	if wantDisplayValue != gotDisplayValue {
		t.Errorf("wantDisplayValue: %d, gotDisplayValue: %d", wantDisplayValue, gotDisplayValue)
	}
}

func TestSumAll(t *testing.T) {
	tests := []struct {
		name string
		fn   string
		want int
	}{
		{
			name: "testdata",
			fn:   "test_data.txt",
			want: 61229,
		},
		{
			name: "data",
			fn:   "input.txt",
			want: 1096964,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ReadData(tt.fn)
			if err != nil {
				t.Fatal(err)
			}

			got := 0
			for _, d := range data {
				signals := GetOrderedSignals(d)
				got += GetDisplayValue(signals, d.Output)
			}

			if tt.want != got {
				t.Errorf("want: %d, got: %d", tt.want, got)
			}
		})
	}
}

func dumpSignals(s [inputCount]Signal) string {

	var sb strings.Builder

	for i, v := range s {
		sb.WriteString(fmt.Sprintf("%d: %s\n", i, string(v)))
	}

	return sb.String()
}
