package quick

// Sort implements quick sort algorythm
func Sort(a []int) (swaps int, passes int) {
	swaps, passes = quicksort(a, 0, len(a)-1)
	return
}

func quicksort(a []int, low int, high int) (swaps int, passes int) {
	passes++
	if low < high {
		// partitioning:
		pivotIndex, swapsOne := partition(a, low, high)
		// now on left of pivotIndex elemnts lesser than a[pivotIndex], on right - elements greater than a[pivotIndex]
		swaps += swapsOne
		// repeat partitioning for left and right half
		swapsOne, passesOne := quicksort(a, low, pivotIndex-1)
		swaps += swapsOne
		passes += passesOne
		swapsOne, passesOne = quicksort(a, pivotIndex, high)
		swaps += swapsOne
		passes += passesOne
	}
	return
}

// partition follow Hoare method
func partition(a []int, low int, high int) (pivotIndex int, swaps int) {
	// select pivot element: in the middle of segment
	middleInd := (low + high) / 2
	pivot := a[middleInd]
	lowCursor := low
	highCursor := high

	// went to pivot, on left of pivot we have lesser elemnts, on right we have greater elements thann pivot
	for {
		// see on elements from begin and from end to find pair of elemnts that sits in wrong half of segments
		for ; lowCursor <= high && a[lowCursor] < pivot; lowCursor++ {
		}
		for ; highCursor >= low && a[highCursor] > pivot; highCursor-- {
		}
		if lowCursor >= highCursor {
			break
		}
		swaps++
		// found elements that sits in wrong half - swaps they
		a[lowCursor], a[highCursor] = a[highCursor], a[lowCursor]
		lowCursor++
		highCursor--
	}

	pivotIndex = highCursor + 1
	return
}
