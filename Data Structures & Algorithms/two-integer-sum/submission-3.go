func twoSum(nums []int, target int) []int {
	compliment := []int{}
	for _, num := range nums {
		compliment = append(compliment, (target - num))
	}
	// fmt.Println(compliment)
	var indices []int
	if len(indices) < 2 {
		for i := range len(nums) {
			for j := range len(compliment) {
				if nums[i] == compliment[j] && i != j {
					indices = append(indices, j)
					break
				} else {
					continue
				}
			}
		}
	}
	if indices[0] > indices[1] {
		indices[0], indices[1] = indices[1], indices[0]
	}
	// fmt.Println(indices)
	return indices
}
