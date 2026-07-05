func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
	// Check with the last element to begin with
	if target > matrix[len(matrix) - 1][len(matrix[len(matrix) - 1]) - 1] {
		return false
	}
	var start1, start2, end1, end2 int
	// start1 and end1 handle the row dimension
	// start2 and end2 handle the column dimension
	start1 = 0
	end1 = len(matrix) - 1

	// Check for the existence of the target in the range specified by the rows
	// Use binary search to narrow down the range

	for start1 <= end1 {
		mid1 := int((start1 + end1)/2)
		// if the target is in the mid range of the 2D matrix
		if target >= matrix[mid1][0] && target <= matrix[mid1][len(matrix[mid1]) - 1] {
			// Run an inner binary search loop here to pin point the target
			temp := matrix[mid1]
			start2 = 0
			end2 = len(temp) - 1
			
			for start2 <= end2 {
				mid2 := int((start2 + end2)/2)
				if target == temp[mid2] {
					return true
				} else if target < temp[mid2] {
					end2 = mid2 - 1
				} else if target > temp[mid2] {
					start2 = mid2 + 1
				}
			}
			return false
		} else if target < matrix[mid1][0] {
			// Change the search range. Making end1 as mid1 - 1
			end1 = mid1 - 1
		} else if target > matrix[mid1][len(matrix[mid1]) - 1] {
			// Change the search range. Make start1 as mid1 + 1
			start1 = mid1 + 1
		}
	}

	return false
}
