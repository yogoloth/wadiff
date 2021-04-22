package main

import (
	"testing"
)

func TestGetPrev(t *testing.T) {
	type test_data struct {
		matrix [][]int
		x      int
		y      int
	}

	test1 := test_data{[][]int{
		{1, 1},
		{1, 2},
	}, 0, 0}
	test2 := test_data{[][]int{
		{2, 2},
		{1, 2},
	}, 1, 0}
	test3 := test_data{[][]int{
		{5, 6},
		{4, 5},
	}, 1, 0}

	tests := []test_data{test1, test2, test3}

	for _, test := range tests {
		if x, y := GetPrev(test.matrix, 1, 1); x != test.x || y != test.y {
			t.Errorf("TestGetPrev %v %d %d return test[%d][%d]=%d\n", test.matrix, 1, 1, x, y, test.matrix[x][y])
		}
	}
}

func TestGetRoute(t *testing.T) {

}
