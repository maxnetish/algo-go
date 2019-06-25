package merge

import "testing"

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
		name        string
		args        args
		wantInserts int
		wantPasses  int
	}{
		{
			name: "Small array",
			args: args{
				a: []int{1, 3, 5, 4, 2, 1, 1, 1, 325, 2},
			},
		},
		{
			name: "Bigger array",
			args: args{
				a: []int{1, 3, 5, 4, 2, 1, 1, 1, 325, 2, 1, 3, 5, 4, 2, 11, 21, 1, 325, 2, 1, 3, 5, 4, 2, 14, 10, 325, 2},
			},
		},
		{
			name: "Empty array",
			args: args{
				a: []int{},
			},
		},
		{
			name: "Single element",
			args: args{
				a: []int{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _, _ := Sort(tt.args.a)
			if !check(result) {
				t.Fail()
			}
		})
	}
}
