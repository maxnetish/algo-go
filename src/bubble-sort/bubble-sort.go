package bubble

// BubbleSorter implements interface Sorter.
var BubbleSorter = struct {
	Sort func([]int)
}{

	Sort: func(a []int) {
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
	},
}
