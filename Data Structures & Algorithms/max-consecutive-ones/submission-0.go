func findMaxConsecutiveOnes(nums []int) int {
	var count, maxCount int
    for i := range len(nums) {
        if nums[i] == 1 {
            count += 1
        } else if nums[i] == 0 {
            if count > maxCount {
                maxCount = count
                count = 0
            } else {
                count = 0
            }
        }
    }
    if count > maxCount {
        maxCount = count
    }
    return maxCount
}
