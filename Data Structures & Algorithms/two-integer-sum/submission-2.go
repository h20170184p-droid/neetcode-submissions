func twoSum(nums []int, target int) []int {
	original := nums
	var compli []int
	for _, val := range nums {
		compli = append(compli, (target - val))
	}
	kvpairs := make(map[int]int)
	original_index := make(map[int]int)

	for index, value := range original {
		kvpairs[value] += 1
		original_index[value] = index
	}

	var i int
	var j int
	for index2, value2 := range compli {
		if count, exists := kvpairs[value2]; exists {
			if value2 == original[index2] {
				if count > 1 {
					i = original_index[value2]
					j = index2
					break
				}
			} else {
				i = original_index[value2]
				j = index2
				break
			}
		}
	}
	if i > j {
		i, j = j, i
	}

	return []int{i, j}
}
