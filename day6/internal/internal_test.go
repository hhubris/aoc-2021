package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReadData(t *testing.T) {
	want := map[int]int{
		1: 1,
		2: 1,
		3: 2,
		4: 1,
	}

	got, err := ReadData("../test_data.txt")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %#v, got %#v", want, got)
	}
}

func TestNewGeneration(t *testing.T) {

	// iv := []int{3, 4, 3, 1, 2}
	iv := map[int]int{
		1: 1,
		2: 1,
		3: 2,
		4: 1,
	}

	wantList := [][]int{
		{2, 3, 2, 0, 1},                                                                //1
		{1, 2, 1, 6, 0, 8},                                                             //2
		{0, 1, 0, 5, 6, 7, 8},                                                          //3
		{6, 0, 6, 4, 5, 6, 7, 8, 8},                                                    //4
		{5, 6, 5, 3, 4, 5, 6, 7, 7, 8},                                                 //5
		{4, 5, 4, 2, 3, 4, 5, 6, 6, 7},                                                 //6
		{3, 4, 3, 1, 2, 3, 4, 5, 5, 6},                                                 //7
		{2, 3, 2, 0, 1, 2, 3, 4, 4, 5},                                                 //8
		{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 8},                                              //9
		{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 7, 8},                                           //10
		{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 7, 8, 8, 8},                                  //11
		{5, 6, 5, 3, 4, 5, 6, 0, 0, 1, 5, 6, 7, 7, 7, 8, 8},                            //12
		{4, 5, 4, 2, 3, 4, 5, 6, 6, 0, 4, 5, 6, 6, 6, 7, 7, 8, 8},                      //13
		{3, 4, 3, 1, 2, 3, 4, 5, 5, 6, 3, 4, 5, 5, 5, 6, 6, 7, 7, 8},                   //14
		{2, 3, 2, 0, 1, 2, 3, 4, 4, 5, 2, 3, 4, 4, 4, 5, 5, 6, 6, 7},                   //15
		{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 1, 2, 3, 3, 3, 4, 4, 5, 5, 6, 8},                //16
		{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 0, 1, 2, 2, 2, 3, 3, 4, 4, 5, 7, 8},             //17
		{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8}, //18
	}

	wantMaps := make([]map[int]int, len(wantList))

	for i, lst := range wantList {
		wantMaps[i] = map[int]int{}
		for _, v := range lst {
			wantMaps[i][v]++
		}
	}

	cg := iv
	for i, want := range wantMaps {
		t.Run(fmt.Sprintf("day %d", i+1), func(t *testing.T) {
			cg = NextGeneration(cg)
			if !reflect.DeepEqual(want, cg) {
				t.Errorf("want: %v, got %v", want, cg)
			}
		})
	}

	cg = iv
	for i := 0; i < 80; i++ {
		cg = NextGeneration(cg)
	}

	tot := CountFish(cg)
	if tot != 5934 {
		t.Errorf("want: 5934, got: %d", tot)
	}
}
