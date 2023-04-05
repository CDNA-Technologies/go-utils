package strings

func AreAnagrams(s1, s2 string) bool {
	// If both strings are of unequal length, they are not anagrams
	if len(s1) != len(s2) {
		return false
	}

	var charCount = make(map[rune]int)

	for _, ch := range s1 {
		charCount[ch]++
	}

	for _, ch := range s2 {
		charCount[ch]--
		if charCount[ch] < 0 {
			return false
		}
	}

	return true
}
