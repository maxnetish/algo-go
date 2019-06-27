package shell

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

func Test_sorter_Sort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		s    *sorter
		args args
	}{
		{
			name: "Empty array",
			s:    &Sorter,
			args: args{
				a: []int{},
			},
		},
		{
			name: "Array len=1",
			s:    &Sorter,
			args: args{
				a: []int{1},
			},
		},
		{
			name: "Array len=2",
			s:    &Sorter,
			args: args{
				a: []int{1, 3},
			},
		},
		{
			name: "Array len=3",
			s:    &Sorter,
			args: args{
				a: []int{1, 3, 2},
			},
		},
		{
			name: "Array len=4",
			s:    &Sorter,
			args: args{
				a: []int{1, 3, 2, 4},
			},
		},
		{
			name: "Array len=5",
			s:    &Sorter,
			args: args{
				a: []int{5, 1, 3, 2, 4},
			},
		},
		{
			name: "Array len=10",
			s:    &Sorter,
			args: args{
				a: []int{5, 1, 3, 2, 4, 5, 1, 3, 2, 4},
			},
		},
		{
			name: "Array len=17",
			s:    &Sorter,
			args: args{
				a: []int{5, 1, 3, 2, 4, 5, 1, 3, 2, 4, 2, 4, 5, 1, 3, 2, 4},
			},
		},
		{
			name: "Array len=34",
			s:    &Sorter,
			args: args{
				a: []int{5, 1, 3, 2, 4, 5, 1, 3, 2, 4, 2, 4, 5, 1, 3, 2, 4, 5, 1, 3, 2, 4, 5, 1, 3, 2, 4, 2, 4, 5, 1, 3, 2, 4},
			},
		},
		{
			name: "Array len=68",
			s:    &Sorter,
			args: args{
				a: []int{5, 1, 3, 2, 4, 5, 1, 3, 2, 4, 2, 4, 5, 1, 3, 2, 4, 5, 1, 3, 2, 4, 5, 1, 3, 2, 4, 2, 4, 5, 1, 3, 2, 4, 5, 1, 3, 2, 4, 5, 1, 3, 2, 4, 2, 4, 5, 1, 3, 2, 4, 5, 1, 3, 2, 4, 5, 1, 3, 2, 4, 2, 4, 5, 1, 3, 2, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Sort(tt.args.a)
			if !check(tt.args.a) {
				t.Fail()
			}
		})
	}
}
