package sort

import "golang.org/x/exp/constraints"

// Bubble implements bubble sort algorithm. It works by iterating over elements
// over and over until collection becomes sorted. It finds out of order adjacent
// elements on each iteration and swap them. As a result, out of order elements
// eventually bubble up to their valid positions.
func Bubble[T constraints.Ordered](values []T) {
	valuesCount := len(values)

	if valuesCount < 1 {
		return
	}

	var isSorted bool

	for !isSorted {
		isSorted = true
		for i := 1; i < valuesCount; i++ {
			if values[i] < values[i-1] {
				values[i], values[i-1] = values[i-1], values[i]
				isSorted = false
			}
		}
	}
}
