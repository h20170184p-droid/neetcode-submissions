func uniquePaths(m int, n int) int {
    if m == 1 || n == 1 {
		return 1
	} else if m < 1 || n < 1 {
		return 0
	}
	matrix := make([][]int, m)
	for i := range len(matrix) {
		matrix[i] = make([]int, n)
	}
	for i := range len(matrix) {
		matrix[i][0] = 1
	}
	for i := range matrix[0] {
		matrix[0][i] = 1
	}
	// return matrix
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}
	return matrix[m-1][n-1]
}
