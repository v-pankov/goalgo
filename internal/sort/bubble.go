package sort

import (
	"golang.org/x/exp/constraints"
)

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

// Bubble1 is an optimized version of Bubble which skips iterating over elements
// which are already put in proper position.
func Bubble1[T constraints.Ordered](values []T) {
	valuesCount := len(values)

	if valuesCount < 1 {
		return
	}

	// 'i' stands for count of elements which were put at the end.
	// The greatest value is put to the end on each iteration.
	// Skip this value on next iteration.
	for i := 0; i < valuesCount-1; i++ {
		isSorted := true
		for j := 1; j < valuesCount-i; j++ {
			if values[j] < values[j-1] {
				values[j], values[j-1] = values[j-1], values[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
}

// Bubble2 is an optimized version of Bubble1 which does not process elements
// after previous iteration last swapped index.
func Bubble2[T constraints.Ordered](values []T) {
	valuesCount := len(values)

	if valuesCount < 1 {
		return
	}

	for end := valuesCount; end >= 0; {
		lastSwappedIdx := -1

		for j := 1; j < end; j++ {
			if values[j] < values[j-1] {
				values[j], values[j-1] = values[j-1], values[j]
				lastSwappedIdx = j
			}
		}

		if lastSwappedIdx >= 0 {
			end = lastSwappedIdx
		} else {
			break
		}
	}
}

// Bubble2 is like Bubble2 but changes sort direction on each iteration (a-la Cocktail sort)
func Bubble3[T constraints.Ordered](values []T) {
	valuesCount := len(values)

	if valuesCount < 1 {
		return
	}

	var (
		dirDown  bool
		leftEnd  int
		rightEnd = valuesCount
	)

	for {
		shakerFn := shakeUp[T]
		if dirDown {
			shakerFn = shakeDown[T]
		}

		done, newEnd := shakerFn(
			leftEnd, rightEnd, values,
		)

		if done {
			break
		}

		if dirDown {
			leftEnd = newEnd
		} else {
			rightEnd = newEnd
		}

		dirDown = !dirDown
	}
}

func shakeUp[T constraints.Ordered](
	leftEnd, rightEnd int, values []T,
) (
	bool,
	int,
) {
	lastSwappedIndex := -1
	for i := leftEnd + 1; i < rightEnd; i++ {
		if values[i] < values[i-1] {
			values[i], values[i-1] = values[i-1], values[i]
			lastSwappedIndex = i
		}
	}
	return lastSwappedIndex < 0, lastSwappedIndex
}

func shakeDown[T constraints.Ordered](
	leftEnd, rightEnd int, values []T,
) (
	bool,
	int,
) {
	lastSwappedIndex := -1
	for i := rightEnd - 1; i > leftEnd; i-- {
		if values[i] < values[i-1] {
			values[i], values[i-1] = values[i-1], values[i]
			lastSwappedIndex = i
		}
	}
	return lastSwappedIndex < 0, lastSwappedIndex
}
