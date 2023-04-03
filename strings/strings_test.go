package strings

import (
	"fmt"
	"testing"
)

func TestAreRealAnagramss(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"", "", true},
		{"a", "a", true},
		{"aa", "aa", true},
		{"abcd", "dcba", true},
		{"anagram", "nag a ram", false},
		{"listen", "silent", true},
		{"abcdefg", "gfedcba", true},
		{"abc", "abcd", false},
		{"abcd", "dcbb", false},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestAreRealAnagrams(%v, %v)", input.s1, input.s2), func(t *testing.T) {
			if got := AreRealAnagrams(input.s1, input.s2); got != input.want {
				t.Errorf("AreRealAnagrams() = %v, want %v", got, input.want)
			}
		})
	}
}

func TestAreAnagramss(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"", "", true},
		{"a", "a", true},
		{"aa", "aa", true},
		{"abcd", "dcba", true},
		{"anagram", "nag a ram", true},
		{"listen", "silent", true},
		{"lllllagfl--", "lall-glllf", true},
		{"abc", "abcd", false},
		{"abcd dabc", "abcd", false},
		{":-;][=+", "+;=:", true},
		{":-asds;][=+", "+;=:", false},
		{":-;][=+", "+asds;=:", false},
		{":-;][asdf32=+", "+asds;=:", false},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestAreAnagrams(%v, %v)", input.s1, input.s2), func(t *testing.T) {
			if got := AreAnagrams(input.s1, input.s2); got != input.want {
				t.Errorf("TestAreAnagramss() = %v, want %v", got, input.want)
			}
		})
	}
}
