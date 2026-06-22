func removeElement(nums []int, val int) int {

	var mover int = 0
	for i := range len(nums) {
		if nums[i] != val {
			nums[mover] = nums[i]
			mover += 1
		}
	}
	return mover
}
