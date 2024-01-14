package sort

import (
	"golang.org/x/exp/constraints"
)

// Insertion implements insertion sort algorithm. It keeps track of sorted and
// unsorted subsets of collection. On each iteration, it moves first element of
// unsorted set to the left through the elements in sorted set by swapping them
// until position is found.
func Insertion[T constraints.Ordered](values []T) {
	valuesCount := len(values)

	if valuesCount < 2 {
		return
	}

	// 'i' stands for first element in unsorted set
	for i := 1; i < valuesCount; i++ {
		// Move first element of unsorted set to the left until it's position
		// is found.
		for j := i; j > 0 && values[j] < values[j-1]; j-- {
			values[j], values[j-1] = values[j-1], values[j]
		}
	}
}
