func hasDuplicate(nums []int) bool {
	var found bool = false
	len_slice := make(map[int]int)
	// fmt.Println("The empty checker slice is: ", len_slice)
	for _, value := range nums {

		// Just increment the value at index of choice by 1
		len_slice[value] += 1
		if len_slice[value] > 1 {
			// fmt.Println("Found duplicate!")
			found = true
			break
		}
	}
	return found
}