package strings

import (
	"fmt"
	"testing"
)

func TestAreAnagrams(t *testing.T) {
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
		{":-;][=+", "+;=:", false},
		{":-asds;][=+", "+a;s[-d=:s]", true},
		{":-asds;][=+", "+;=:", false},
		{":-;][=+", "+asds;=:", false},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestAreAnagrams(%v, %v)", input.s1, input.s2), func(t *testing.T) {
			if got := AreAnagrams(input.s1, input.s2); got != input.want {
				t.Errorf("AreAnagrams() = %v, want %v", got, input.want)
			}
		})
	}
}
