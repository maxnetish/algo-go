package merge

// MergeSorter implements Sorter interface
var MergeSorter = struct {
	Sort func([]int)
}{
	Sort: sort,
}

// Sort implements merge sort algorythme (by John von Neumann)
func sort(a []int) {
	mergesortIterative(a)
	return
}

// algorythme without recursion
// can return new slice or a
func mergesortIterative(a []int) {
	aLen := len(a)
	shouldSwap := false
	// slice for buffering intermediate results
	// use O(n) of mem
	intermediate := make([]int, aLen)
	var mergedTo int
	for segmentLen := 1; segmentLen < aLen; segmentLen *= 2 {
		// break a into sublists
		// segmentLen is max len of current sublist
		for left := 0; left < aLen-segmentLen; left += 2 * segmentLen {
			// each time we take two slices: [left:mid] and [mid:right] and merge them into
			// intermediate slice
			mid := left + segmentLen
			right := left + 2*segmentLen
			if right > aLen {
				// to correct merge remaining tail
				right = aLen
			}
			merge(a[left:mid], a[mid:right], intermediate[left:right])
			mergedTo = right
		}
		// copy tail to intermediate else tail will be incorrect
		copy(intermediate[mergedTo:], a[mergedTo:])
		// We use two slices in turns to not to allocate slice in every pass
		// shouldSwap sets if we need copy from a to intermediate else caller will get ref on not actual slice
		shouldSwap = !shouldSwap
		a, intermediate = intermediate, a
	}
	if shouldSwap {
		copy(intermediate, a)
	}
	return
}

// left and right - sorted arrays, merges into result
func merge(left []int, right []int, result []int) {
	cursorLeft := 0
	cursorRight := 0
	cursorResult := 0
	lenLeft := len(left)
	lenRight := len(right)
	for cursorLeft < lenLeft || cursorRight < lenRight {
		if cursorLeft >= lenLeft {
			// items only in right part
			result[cursorResult] = right[cursorRight]
			cursorResult++
			cursorRight++
			continue
		}
		if cursorRight >= lenRight {
			// items remain only in left part
			result[cursorResult] = left[cursorLeft]
			cursorResult++
			cursorLeft++
			continue
		}
		// select lesser item and insert into result
		if left[cursorLeft] <= right[cursorRight] {
			result[cursorResult] = left[cursorLeft]
			cursorResult++
			cursorLeft++
		} else {
			result[cursorResult] = right[cursorRight]
			cursorResult++
			cursorRight++
		}
	}
	return
}
