package shell

type sorter struct{}

var Sorter = sorter{}

func (s *sorter) String() string {
	return "Shell sort (Donald L. Shell sort)"
}

func (s *sorter) Slow() bool {
	return true
}

// Sort implements Shell sort algorythm
func (s *sorter) Sort(a []int) {
	shellSort(a)
	return
}

func shellSort(a []int) {
	aLen := len(a)
	if aLen < 2 {
		return
	}

	distances := getDistances(aLen)

	for distanceIndex := len(distances) - 1; distanceIndex > -1; distanceIndex-- {
		insertSort(a, distances[distanceIndex])
	}
}

func insertSort(a []int, increment int) {
	aLen := len(a)
	aLenWithoutOne := aLen - 1
	for startInd := 0; startInd < increment; startInd++ {
		for ind := startInd; ind < aLenWithoutOne; ind += increment {
			passInd := ind + increment
			if passInd > aLenWithoutOne {
				passInd = aLenWithoutOne
			}
			for ; (passInd - increment) > -1; passInd -= increment {
				prevPassInd := passInd - increment
				if a[prevPassInd] > a[passInd] {
					a[prevPassInd], a[passInd] = a[passInd], a[prevPassInd]
				}
			}
		}
	}
}

func getDistances(aLength int) []int {
	if aLength < 2 {
		return []int{}
	}

	ind := len(distancesByPanchenko) - 1
	// ind > 0 because we will return empty slice for such case, see higher
	for ; ind > 0 && distancesByPanchenko[ind] >= aLength; ind-- {
	}

	return distancesByPanchenko[:ind+1]
}

// magic numbers
// sequence https://oeis.org/A055875
// Ivan Panchenko
//
var distancesByPanchenko = []int{
	1,
	2,
	19,
	103,
	311,
	691,
	1321,
	2309,
	3671,
	5519,
	7919,
	10957,
	14753,
	19403,
	24809,
	31319,
	38873,
	47657,
	57559,
	69031,
	81799,
	96137,
	112291,
	130073,
	149717,
	171529,
	195043,
	220861,
	248851,
	279431,
	312583,
	347707,
	386093,
	427169,
	470933,
	517553,
	567871,
	620531,
	677539,
	737203,
	800573,
	867677,
	938533,
}
