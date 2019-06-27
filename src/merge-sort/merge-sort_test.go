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
			name: "Small array 2",
			args: args{
				a: []int{1, 3, 5, 4, 2, 1, 1, 1, 325, 2, 200},
			},
		},
		{
			name: "Small array 2",
			args: args{
				a: []int{1, 3, 5, 4, 2, 1, 1, 1, 0, 2, 200, 199},
			},
		},
		{
			name: "Small array 3",
			args: args{
				a: []int{1, 3, 5, 4, 2, 1, 1, 1, 0, 2, 200, 199, 198},
			},
		},
		{
			name: "Small array 4",
			args: args{
				a: []int{1, 3, 5, 4, 2, 1, 1, 1, 0, 2, 200, 199, 198, 197},
			},
		},
		{
			name: "Small array 5",
			args: args{
				a: []int{196, 1, 3, 5, 4, 2, 1, 1, 1, 0, 2, 200, 199, 198, 197},
			},
		},
		{
			name: "Small array 6",
			args: args{
				a: []int{196, 195, 1, 3, 5, 4, 2, 1, 1, 1, 0, 2, 200, 199, 198, 197},
			},
		},
		{
			name: "Small array 7",
			args: args{
				a: []int{196, 195, 194, 1, 3, 5, 4, 2, 1, 1, 1, 0, 2, 200, 199, 198, 197},
			},
		},
		{
			name: "Small array 8",
			args: args{
				a: []int{196, 195, 194, 193, 1, 3, 5, 4, 2, 1, 1, 1, 0, 2, 200, 199, 198, 197},
			},
		},
		{
			name: "Small array 9",
			args: args{
				a: []int{196, 195, 194, 193, 192, 1, 3, 5, 4, 2, 1, 1, 1, 0, 2, 200, 199, 198, 197},
			},
		},
		{
			name: "Small array 10",
			args: args{
				a: []int{196, 195, 194, 193, 192, 191, 1, 3, 5, 4, 2, 1, 1, 1, 0, 2, 200, 199, 198, 197},
			},
		},
		{
			name: "Bigger array",
			args: args{
				a: []int{1, 3, 5, 4, 2, 1, 1, 1, 325, 2, 1, 3, 5, 4, 2, 11, 21, 1, 325, 2, 1, 3, 5, 4, 2, 14, 10, 325, 2, 0},
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
			Sorter.Sort(tt.args.a)
			if !check(tt.args.a) {
				t.Fail()
			}
		})
	}
}
