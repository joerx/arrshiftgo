package zeroshift

import "testing"

var input = []int{8, -9, 0, 4, 0, 0, 5, 0, -1}
var expected = []int{8, -9, -1, 4, 5, 0, 0, 0, 0}

func TestNaiveShift(t *testing.T) {
	actual := NaiveShift(input)
	if !validate(actual) {
		t.Error("Found non-zero following a zero in result")
	}
}

func TestFastShift(t *testing.T) {
	actual := FastShift(input)
	if !validate(actual) {
		t.Error("Found non-zero following a zero in result")
	}
}

func sumSlice(sl []int) (r int) {
	for _, n := range sl {
		r += n
	}
	return
}

// validation: any zero in result is only followed by another zero or end of array
func validate(actual []int) bool {
	for i := 1; i < len(actual); i++ {
		if actual[i-1] == 0 && actual[i] != 0 {
			return false
		}
	}
	return true
}
