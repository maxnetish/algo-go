package quick

import (
	"testing"

	"github.com/maxnetish/algo-go/src/utils"
)

func check(a []int) bool {
	lenA := len(a)
	if lenA == 0 || lenA == 1 {
		return true
	}
	var prev int
	for ind := 1; ind < lenA; ind++ {
		prev = ind - 1
		if a[prev] > a[ind] {
			return false
		}
	}
	return true
}

func TestSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sort big array",
			args: args{
				// a: []int{0, 0, 1, 1, 3, 2, 3, 4, 5, 4},
				a: utils.MakeBigArray(utils.MakeBigArrayParams{NumberOfElements: 1000, To: 20}),
			},
		},
		{
			name: "Sort tiny array",
			args: args{
				a: []int{1, 0, 1},
				// a: utils.MakeBigArray(utils.MakeBigArrayParams{NumberOfElements: 3, To: 2}),
			},
		},
		{
			name: "Array single element",
			args: args{
				a: []int{123},
				// a: utils.MakeBigArray(utils.MakeBigArrayParams{NumberOfElements: 1, To: 2}),
			},
		},
		{
			name: "Empty array",
			args: args{
				a: make([]int, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sorter.Sort(tt.args.a)
			if !check(tt.args.a) {
				t.Fail()
			}
		})
	}
}
