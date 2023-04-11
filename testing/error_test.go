package testing

import (
	"fmt"
	"testing"
)

func TestIsErrorEqual(t *testing.T) {
	tests := []struct {
		a    error
		b    error
		want bool
	}{
		{
			a:    nil,
			b:    nil,
			want: true,
		},
		{
			a:    nil,
			b:    fmt.Errorf("Error"),
			want: false,
		},
		{
			a:    fmt.Errorf("Error"),
			b:    nil,
			want: false,
		},
		{
			a:    fmt.Errorf("Error 1"),
			b:    fmt.Errorf("Error 2"),
			want: false,
		},
		{
			a:    fmt.Errorf("Error"),
			b:    fmt.Errorf("Error"),
			want: true,
		},
	}
	for _, input := range tests {
		t.Run("IsErrorEqual", func(t *testing.T) {
			if got := IsErrorEqual(input.a, input.b); got != input.want {
				t.Errorf("IsErrorEqual() = %v, want %v", got, input.want)
			}
		})
	}
}
