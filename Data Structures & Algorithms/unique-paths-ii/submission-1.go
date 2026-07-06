func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
        return 0
    }

	matrix := make([][]int, m)
	for i := range len(matrix) {
		matrix[i] = make([]int, n)
	}

	matrix[m - 1][n - 1] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >=0; j-- {
			if i == m - 1 && j == n - 1 {
                continue
            }
			if obstacleGrid[i][j] == 1 {
				matrix[i][j] = 0
			} else {
				if i == (m - 1) {
					matrix[i][j] = matrix[i][j + 1]
				} else if j == (n - 1){
					matrix[i][j] = matrix[i + 1][j]
				} else {
					matrix[i][j] = matrix[i][j + 1] + matrix[i + 1][j]
				} 
			}
		}
	}

	return matrix[0][0]
}
