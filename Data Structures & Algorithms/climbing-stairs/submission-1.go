func climbStairs(n int) int {
    memo := make(map[int]int)
    var helper func(int) int
    helper = func(n int) int {
        if n <= 2 {
            return n
        }
        if val, ok := memo[n]; ok {
            return val
        }
        memo[n] = helper(n-1) + helper(n-2)
        return memo[n]
    }
    return helper(n)
}
