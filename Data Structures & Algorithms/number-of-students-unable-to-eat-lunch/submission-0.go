func countStudents(students []int, sandwiches []int) int {
	counter := 0
    for {
		if len(sandwiches) == 0 {
            break
        }
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			counter = 0
		} else {
			mover := students[0]
			students = append(students[1:], mover)
			counter += 1
			if counter == len(sandwiches) {
				break
			}
		}
		// students = append(students, mover)
	}
	return counter
}