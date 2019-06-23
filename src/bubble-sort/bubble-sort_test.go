package bubble

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
			name: "Check sort - small array",
			args: args{
				a: utils.MakeBigArray(
					utils.MakeBigArrayParams{NumberOfElements: 6, To: 50},
				),
			},
		},
		{
			name: "Check sort - big array",
			args: args{
				a: utils.MakeBigArray(
					utils.MakeBigArrayParams{NumberOfElements: 10000, To: 100},
				),
			},
		},
		{
			name: "Check sort - empty array",
			args: args{
				a: []int{},
			},
		},
		{
			name: "Check sort - array of single element",
			args: args{
				a: []int{123456},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.a)
			if !check(tt.args.a) {
				t.Fail()
			}
		})
	}
}
