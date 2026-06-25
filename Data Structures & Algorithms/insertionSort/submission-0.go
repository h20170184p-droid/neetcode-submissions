// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
	output := make([][]Pair, 0)
	for i := range len(pairs) {
		for j := i; j > 0; {
			if pairs[j].Key < pairs[j -1].Key {
				pairs[j], pairs[j -1] = pairs[j -1], pairs[j]
			}
			j -= 1
		}
		temp := make([]Pair, len(pairs))
		copy(temp, pairs)
		output = append(output, temp)
	}
	return output
}
