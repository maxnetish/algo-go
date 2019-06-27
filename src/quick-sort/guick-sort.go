package quick

type sorter struct{}

// Sorter singletone implements Sorter interface
var Sorter = sorter{}

func (s *sorter) String() string {
	return "Quick sort (Charles Antony Richard Hoare sort)"
}

func (s *sorter) Slow() bool {
	return false
}

// Sort implements quick sort algorythm
func (s *sorter) Sort(a []int) {
	quicksort(a, 0, len(a)-1)
	return
}

func quicksort(a []int, low int, high int) {
	if low < high {
		// partitioning:
		pivotIndex := partition(a, low, high)
		// now on left of pivotIndex elemnts lesser than a[pivotIndex], on right - elements greater than a[pivotIndex]
		// repeat partitioning for left and right half
		quicksort(a, low, pivotIndex-1)
		quicksort(a, pivotIndex, high)
	}
	return
}

// partition follow Hoare method
func partition(a []int, low int, high int) (pivotIndex int) {
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
		// found elements that sits in wrong half - swaps they
		a[lowCursor], a[highCursor] = a[highCursor], a[lowCursor]
		lowCursor++
		highCursor--
	}

	pivotIndex = highCursor + 1
	return
}
