func getConcatenation(nums []int) []int {
    ans := make([]int, (2*len(nums)))
	n := len(nums)
	for i := range n {
		ans[i] = nums[i]
		ans[i + n] = nums[i]
	}
	return ans
}
