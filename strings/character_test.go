package strings

import (
	"fmt"
	"testing"
)

func TestIsUpper(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: true},
		{args: args{ch: 'R'}, want: true},
		{args: args{ch: 'F'}, want: true},
		{args: args{ch: 'Q'}, want: true},
		{args: args{ch: 'Z'}, want: true},
		{args: args{ch: 'a'}, want: false},
		{args: args{ch: 'g'}, want: false},
		{args: args{ch: '0'}, want: false},
		{args: args{ch: '5'}, want: false},
		{args: args{ch: '9'}, want: false},
		{args: args{ch: 'z'}, want: false},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestIsUpper(%v)", input.args.ch), func(t *testing.T) {
			if got := IsUpper(input.args.ch); got != input.want {
				t.Errorf("IsUpper(%v) = %v, want %v", input.args.ch, got, input.want)
			}
		})
	}
}

func TestIsLower(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: false},
		{args: args{ch: 'R'}, want: false},
		{args: args{ch: 'F'}, want: false},
		{args: args{ch: 'Q'}, want: false},
		{args: args{ch: 'Z'}, want: false},
		{args: args{ch: 'a'}, want: true},
		{args: args{ch: 'g'}, want: true},
		{args: args{ch: '0'}, want: false},
		{args: args{ch: '5'}, want: false},
		{args: args{ch: '9'}, want: false},
		{args: args{ch: 'z'}, want: true},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestIsLower(%v)", input.args.ch), func(t *testing.T) {
			if got := IsLower(input.args.ch); got != input.want {
				t.Errorf("IsLower(%v) = %v, want %v", input.args.ch, got, input.want)
			}
		})
	}
}

func TestIsLetter(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: true},
		{args: args{ch: 'R'}, want: true},
		{args: args{ch: 'F'}, want: true},
		{args: args{ch: 'Q'}, want: true},
		{args: args{ch: 'Z'}, want: true},
		{args: args{ch: 'a'}, want: true},
		{args: args{ch: 'g'}, want: true},
		{args: args{ch: '0'}, want: false},
		{args: args{ch: '5'}, want: false},
		{args: args{ch: '9'}, want: false},
		{args: args{ch: 'z'}, want: true},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestIsLetter(%v)", input.args.ch), func(t *testing.T) {
			if got := IsLetter(input.args.ch); got != input.want {
				t.Errorf("TestIsLetter(%v) = %v, want %v", input.args.ch, got, input.want)
			}
		})
	}
}

func TestIsDigit(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: false},
		{args: args{ch: 'R'}, want: false},
		{args: args{ch: 'F'}, want: false},
		{args: args{ch: 'Q'}, want: false},
		{args: args{ch: 'Z'}, want: false},
		{args: args{ch: 'a'}, want: false},
		{args: args{ch: 'g'}, want: false},
		{args: args{ch: '0'}, want: true},
		{args: args{ch: '5'}, want: true},
		{args: args{ch: '9'}, want: true},
		{args: args{ch: 'z'}, want: false},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestIsDigit(%v)", input.args.ch), func(t *testing.T) {
			if got := IsDigit(input.args.ch); got != input.want {
				t.Errorf("TestIsDigit(%v) = %v, want %v", input.args.ch, got, input.want)
			}
		})
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: true},
		{args: args{ch: 'R'}, want: true},
		{args: args{ch: 'F'}, want: true},
		{args: args{ch: 'Q'}, want: true},
		{args: args{ch: 'Z'}, want: true},
		{args: args{ch: 'a'}, want: true},
		{args: args{ch: 'g'}, want: true},
		{args: args{ch: '0'}, want: true},
		{args: args{ch: '5'}, want: true},
		{args: args{ch: '9'}, want: true},
		{args: args{ch: 'z'}, want: true},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestIsAlphaNumeric(%v)", input.args.ch), func(t *testing.T) {
			if got := IsAlphaNumeric(input.args.ch); got != input.want {
				t.Errorf("TestIsAlphaNumeric(%v) = %v, want %v", input.args.ch, got, input.want)
			}
		})
	}
}

func TestIsVowel(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: true},
		{args: args{ch: 'E'}, want: true},
		{args: args{ch: 'R'}, want: false},
		{args: args{ch: 'F'}, want: false},
		{args: args{ch: 'Q'}, want: false},
		{args: args{ch: 'Z'}, want: false},
		{args: args{ch: 'a'}, want: true},
		{args: args{ch: 'u'}, want: true},
		{args: args{ch: 'g'}, want: false},
		{args: args{ch: '0'}, want: false},
		{args: args{ch: '5'}, want: false},
		{args: args{ch: '9'}, want: false},
		{args: args{ch: 'z'}, want: false},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestIsVowel(%v)", input.args.ch), func(t *testing.T) {
			if got := IsVowel(input.args.ch); got != input.want {
				t.Errorf("TestIsVowel(%v) = %v, want %v", input.args.ch, got, input.want)
			}
		})
	}
}

func TestIsConsonant(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: false},
		{args: args{ch: 'E'}, want: false},
		{args: args{ch: 'R'}, want: true},
		{args: args{ch: 'F'}, want: true},
		{args: args{ch: 'Q'}, want: true},
		{args: args{ch: 'Z'}, want: true},
		{args: args{ch: 'a'}, want: false},
		{args: args{ch: 'u'}, want: false},
		{args: args{ch: 'g'}, want: true},
		{args: args{ch: '0'}, want: false},
		{args: args{ch: '5'}, want: false},
		{args: args{ch: '9'}, want: false},
		{args: args{ch: 'z'}, want: true},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TestIsConsonant(%v)", input.args.ch), func(t *testing.T) {
			if got := IsConsonant(input.args.ch); got != input.want {
				t.Errorf("TestIsConsonant(%v) = %v, want %v", input.args.ch, got, input.want)
			}
		})
	}
}

func TesinputoLower(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want rune
	}{
		{args: args{ch: 'A'}, want: 'a'},
		{args: args{ch: 'E'}, want: 'e'},
		{args: args{ch: 'R'}, want: 'r'},
		{args: args{ch: 'F'}, want: 'f'},
		{args: args{ch: 'Q'}, want: 'q'},
		{args: args{ch: 'Z'}, want: 'z'},
		{args: args{ch: 'a'}, want: 'a'},
		{args: args{ch: 'u'}, want: 'u'},
		{args: args{ch: 'g'}, want: 'g'},
		{args: args{ch: '0'}, want: '0'},
		{args: args{ch: '5'}, want: '5'},
		{args: args{ch: '9'}, want: '9'},
		{args: args{ch: 'z'}, want: 'z'},
		{args: args{ch: '['}, want: '['},
		{args: args{ch: '#'}, want: '#'},
		{args: args{ch: '='}, want: '='},
		{args: args{ch: '\''}, want: '\''},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TesinputoLower(%v)", input.args.ch), func(t *testing.T) {
			if got := ToLower(input.args.ch); got != input.want {
				t.Errorf("TesinputoLower(%v) = %v, want %v", input.args.ch, got, input.want)
			}
		})
	}
}

func TesinputoUpper(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want rune
	}{
		{args: args{ch: 'A'}, want: 'A'},
		{args: args{ch: 'E'}, want: 'E'},
		{args: args{ch: 'R'}, want: 'R'},
		{args: args{ch: 'F'}, want: 'F'},
		{args: args{ch: 'Q'}, want: 'Q'},
		{args: args{ch: 'Z'}, want: 'Z'},
		{args: args{ch: 'a'}, want: 'A'},
		{args: args{ch: 'u'}, want: 'U'},
		{args: args{ch: 'g'}, want: 'G'},
		{args: args{ch: '0'}, want: '0'},
		{args: args{ch: '5'}, want: '5'},
		{args: args{ch: '9'}, want: '9'},
		{args: args{ch: 'z'}, want: 'Z'},
		{args: args{ch: '['}, want: '['},
		{args: args{ch: '#'}, want: '#'},
		{args: args{ch: '='}, want: '='},
		{args: args{ch: '\''}, want: '\''},
	}
	for _, input := range tests {
		t.Run(fmt.Sprintf("TesinputoUpper(%v)", input.args.ch), func(t *testing.T) {
			if got := ToUpper(input.args.ch); got != input.want {
				t.Errorf("TesinputoUpper(%v) = %v, want %v", input.args.ch, got, input.want)
			}
		})
	}
}
