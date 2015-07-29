package sortp

import (
	"sort"
)

// The bubble sort algorithm. It is especially useful for almost sorted list.
// Bubble sort is a stable sort algorithm.
func Bubble(data sort.Interface) {
	BubbleF(data.Len(), data.Less, data.Swap)
}

// Similar to Bubble but using closures other than sort.Interface.
func BubbleF(Len int, Less func(i, j int) bool, Swap func(i, j int)) {
	for Len > 0 {
		m := 0
		for i := 1; i < Len; i++ {
			if Less(i, i-1) {
				Swap(i, i-1)
				m = i
			}
		}
		Len = m
	}
}
