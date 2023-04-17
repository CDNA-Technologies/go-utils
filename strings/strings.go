package strings

/**
	Checks if two strings are perfect anagrams of each other using character count comparison.
	Anagrams are formed by rearranging the letters of one string using all the original letters exactly once.

	Example: "listen" and "silent" are perfect anagrams of each other, while "now silent" and "listennow" are not.
**/
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
