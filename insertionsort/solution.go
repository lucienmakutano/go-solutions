package insertionsort

// InsertionSort an implementation the Insertion Sort algorithm
func InsertionSort(items []int) {
	for idx := 1; idx < len(items); idx = idx + 1 {
		currentItem := items[idx]
		loopVar := idx - 1

		for loopVar >= 0 && items[loopVar] > currentItem {
			items[loopVar+1] = items[loopVar]
			loopVar--
		}

		items[loopVar+1] = currentItem
	}
}
