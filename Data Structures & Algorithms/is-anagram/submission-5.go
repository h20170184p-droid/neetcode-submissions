func isAnagram(s string, t string) bool {
	anagram := false
	barr1 := []byte(s)
	barr2 := []byte(t)
	store1 := make(map[byte]int)
	store2 := make(map[byte]int)
	for _, value := range barr1 {
		store1[value] += 1
	}
	for _, value2 := range barr2 {
		store2[value2] += 1
	}

	if len(barr1) == len(barr2) {
		if len(store1) == len(store2) {
			anagram = true
			for key, val := range store1 {
				if store2[key] != val {
					anagram = false
					break

				}
			}
		}
	}

	return anagram

}
