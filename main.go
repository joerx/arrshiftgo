package zeroshift

import "fmt"

// NaiveShift is a naive implementation with O(n^2)
func NaiveShift(arr []int) []int {
	work := make([]int, len(arr))
	copy(work, arr)

	for i := 0; i < len(work); i++ {
		fmt.Println(work)
		if work[i] == 0 {
			for j := i; j < len(work); j++ {
				if work[j] != 0 {
					work[i], work[j] = work[j], work[i]
				}
			}
		}
	}

	return work
}

// FastShift is a faster implementation only needing O(n)
func FastShift(arr []int) []int {
	work := make([]int, len(arr))
	copy(work, arr)

	sp := 0
	ep := len(work) - 1

	for sp != ep {
		if work[ep] == 0 {
			ep--
		}
		if work[sp] != 0 {
			sp++
		}
		if work[sp] == 0 && work[ep] != 0 {
			work[sp], work[ep] = work[ep], work[sp]
		}
	}

	return work
}
