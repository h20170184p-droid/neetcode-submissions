func longestCommonSubsequence(text1 string, text2 string) int {
    m := len(text1)
    n := len(text2)

    // Create a 2D DP matrix initialized to 0
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    // Build the grid step-by-step
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            // If characters match, it's 1 + the best result from the previous diagonal
            if text1[i-1] == text2[j-1] {
                dp[i][j] = 1 + dp[i-1][j-1]
            } else {
                // If they don't match, take the best result from either skipping 
                // a character in text1 or skipping a character in text2
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }

    // The bottom-right corner holds the absolute best answer
    return dp[m][n]
}