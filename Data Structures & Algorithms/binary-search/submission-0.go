func search(nums []int, target int) int {
	// Binary search application.
	// Enter the number to be searched for
	// var sear int
	// fmt.Println("Enter the number to be searched for: ", sear)
	// fmt.Scan(&sear)
	k := nums
	l := len(k)
	// Checking the first, last and middle elements first
	if target == k[0] {
		// fmt.Println("Found the element at index: 0")
		return 0
	}
	if target == k[l-1] {
		// fmt.Println("Found the element at index: ", (l - 1))
		return l - 1
	}
	if target == k[int((l-1)/2)] {
		// fmt.Println("Found the element at index: ", int((l-1)/2))
		return int((l - 1) / 2)
	}

	var index []int
	for j := range len(k) {
		index = append(index, j)
	}
	fmt.Println(index)
	for {
		l = len(k)
		if l == 0 {
			// fmt.Println("Not found")
			return -1
		}
		if k[int((l-1)/2)] < target {
			k = k[(1 + int((l-1)/2)):]
			index = index[(1 + int((l-1)/2)):]
		} else if k[int((l-1)/2)] > target {
			k = k[:int((l-1)/2)]
			index = index[:int((l-1)/2)]
		} else if k[int((l-1)/2)] == target {
			// fmt.Println("The index of the searched integer is: ", index[int((l-1)/2)])
			return index[int((l - 1)/2)]
		}
	}
}
