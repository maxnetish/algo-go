package bubble

type sorter struct{}

// Sorter singletone implements interface Sorter.
var Sorter = sorter{}

func (s *sorter) Slow() bool {
	return true
}

func (s *sorter) String() string {
	return "Bubble sort"
}

// Sort implements bubble sorting algorythme
func (s *sorter) Sort(a []int) {
	aLen := len(a)
	if aLen == 0 || aLen == 1 {
		return
	}
	var prev int
	for passed := false; !passed; {
		passed = true
		for ind := 1; ind < aLen; ind++ {
			prev = ind - 1
			if a[prev] > a[ind] {
				passed = false
				a[prev], a[ind] = a[ind], a[prev]
			}
		}
	}
	return
}
