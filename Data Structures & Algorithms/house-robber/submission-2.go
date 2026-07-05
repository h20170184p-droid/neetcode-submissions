func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }

    // dp[i] will store the absolute max money we can make up to house i
    dp := make([]int, len(nums))

    // Base cases are clear and readable
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1]) // For the first two houses, just pick the bigger one

    // Fill the grid step-by-step
    for i := 2; i < len(nums); i++ {
        // At house i, we have a clear, logical choice:
        // Choice 1: Skip this house -> Take the money from the previous house (dp[i-1])
        // Choice 2: Rob this house -> Take current money + money from 2 houses ago (nums[i] + dp[i-2])
        dp[i] = max(dp[i-1], nums[i]+dp[i-2])
    }

    // The final cell holds the ultimate answer
    return dp[len(nums)-1]
}