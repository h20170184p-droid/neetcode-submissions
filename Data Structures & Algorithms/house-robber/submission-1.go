func rob(nums []int) int {
    var sum1, sum2 int
	for _, val := range nums {
		temp := sum1
		if sum2 + val > sum1 {
			sum1 = sum2 + val
		} else {
			sum1 = sum1
		}
		sum2 = temp
	}
	if sum1 > sum2 {
		return sum1
	}
	return sum2
}