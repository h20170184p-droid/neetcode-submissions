func calPoints(operations []string) int {
	var record []int
	for i := range(len(operations)) {
		content, err := strconv.Atoi(operations[i])
		if err == nil {
			record = append(record, content)
		}
		if operations[i] == "+" {
			val1 := record[len(record)-1]
			val2 := record[len(record)-2]
			record = append(record, val1+val2)
		}
		if operations[i] == "D" {
			d := 2 * record[len(record) - 1]
			record = append(record, d)
		}
		if operations[i] == "C" {
			record = record[:len(record)-1]
		}
	}
	var sum2 int
	for _, k := range record {
		sum2 += k
	}
	return sum2
}
