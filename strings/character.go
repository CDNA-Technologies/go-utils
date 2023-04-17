package strings

var vowels = map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

/**
	Check if the given character is an upper-case English alphabet.
**/
func IsUpper(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

/**
	Check if the given character is an lower-case English alphabet.
**/
func IsLower(ch rune) bool {
	return ch >= 'a' && ch <= 'z'
}

/**
	Check if the given character is a letter - a valid upper-case (or) lower-case English alphabet.
**/
func IsLetter(ch rune) bool {
	return IsUpper(ch) || IsLower(ch)
}

/**
	Check if the given character is a digit from 0-9.
**/
func IsDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

/**
	Check if the given character is alpha numeric - a valid Letter (or) Digit.
**/
func IsAlphaNumeric(ch rune) bool {
	return IsLetter(ch) || IsDigit(ch)
}

/**
	Check if the given character is a vowel. It is case-insensitive.
**/
func IsVowel(ch rune) bool {
	return vowels[ToLower(ch)]
}

/**
	Check if the given character is a consonant. It is case-insensitive.
**/
func IsConsonant(ch rune) bool {
	return !IsVowel(ch) && IsLetter(ch)
}

/**
	Converts the given character to its corresponding lower-case if the given character is upper-case, else returns the same character.
**/
func ToLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + ('a' - 'A')
	}
	return ch
}

/**
	Converts the given character to its corresponding upper-case if the given character is lower-case, else returns the same character.
**/
func ToUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - ('a' - 'A')
	}
	return ch
}
