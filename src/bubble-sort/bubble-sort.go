package bubble

// Sort implements bubble algorythme for []int
func Sort(a []int) (int, int) {
	aLen := len(a)
	swaps := 0
	passes := 0
	if aLen == 0 || aLen == 1 {
		return swaps, passes
	}

	var prev int
	for passed := false; !passed; {
		passes++
		passed = true
		for ind := 1; ind < aLen; ind++ {
			prev = ind - 1
			if a[prev] > a[ind] {
				swaps++
				passed = false
				a[prev], a[ind] = a[ind], a[prev]
			}
		}
	}
	return swaps, passes
}
