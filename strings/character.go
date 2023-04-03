package strings

var vowels = map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func IsUpper(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

func IsLower(ch rune) bool {
	return ch >= 'a' && ch <= 'z'
}

func IsLetter(ch rune) bool {
	return IsUpper(ch) || IsLower(ch)
}

func IsDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func IsAlphaNumeric(ch rune) bool {
	return IsLetter(ch) || IsDigit(ch)
}

func IsVowel(ch rune) bool {
	return vowels[ToLower(ch)]
}

func IsConsonant(ch rune) bool {
	return !IsVowel(ch) && IsLetter(ch)
}

func ToLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + ('a' - 'A')
	}
	return ch
}

func ToUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - ('a' - 'A')
	}
	return ch
}
