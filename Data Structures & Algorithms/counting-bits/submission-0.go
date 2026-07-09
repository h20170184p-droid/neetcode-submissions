func countBits(n int) []int {
	// Crude way to do this is to just run the left shifting check on each number
	// Is there a better way?
	array := make([]int, (n + 1))
	for num := range (n + 1) {
		count := 0
		i := num
		for i > 0 {
			if i & 1 == 1 {
				count += 1
			}
			i = i >> 1
		}
		array[num] = count
	}
	return array
}
