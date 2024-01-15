package sort

import "golang.org/x/exp/constraints"

// Selection implements selection sort algorithm. It works by selecting minimal
// element in unsorted set and swapping it with first element in unsorted set
// thus making it last element of sorted one. The sorted set is initially empty
// and grows by single element each iteration.
func Selection[T constraints.Ordered](values []T) {
	valuesCount := len(values)

	if valuesCount < 2 {
		return
	}

	// 'i' idicates start of unsorted set
	// 'i' does not include last element of collection because it's going to be
	// maximal element at the end.
	for i := 0; i < valuesCount-1; i++ {
		// Assume first element in unsorted set is a miminal one.
		// Remember it's value and index.
		mVal := values[i]
		mIdx := i

		// Iterate over remaining elements in unsorted set.
		for j := i + 1; j < valuesCount; j++ {
			// Update current minimal element and it's index if next element is
			// smaller.
			if values[j] < mVal {
				mVal = values[j]
				mIdx = j
			}
		}

		// Swap first element in unsorted set with minimal element.
		// This extends sorted set by one element.
		if i != mIdx {
			values[i], values[mIdx] = values[mIdx], values[i]
		}
	}
}
