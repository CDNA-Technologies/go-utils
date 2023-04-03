package strings

func AreRealAnagrams(s1, s2 string) bool {
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

func AreAnagrams(s1, s2 string) bool {
	var charCount = make(map[rune]int)

	for _, ch := range s1 {
		if IsAlphaNumeric(ch) {
			charCount[ch]++
		}
	}

	for _, ch := range s2 {
		if !IsAlphaNumeric(ch) {
			continue
		}
		charCount[ch]--
	}

	for k := range charCount {
		if charCount[k] != 0 {
			return false
		}
	}

	return true
}
