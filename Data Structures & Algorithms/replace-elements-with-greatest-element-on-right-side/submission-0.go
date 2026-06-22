func replaceElements(arr []int) []int {
	var max int = 0
	rightArr := arr[1:]
	for i := range (len(arr) - 1) {
		// max = slices.Max(rightArr)
		max = rightArr[0]
		for _, val := range rightArr[1:] {
			if val > max {
				max = val
			}
		}
		arr[i] = max
		if i != (len(arr) - 1) {
			rightArr = rightArr[1:]
		}
	}
	arr [len(arr) - 1] = -1
	return arr
}
