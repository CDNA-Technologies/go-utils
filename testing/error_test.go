package testing

import (
	"fmt"
	"testing"
)

func TestIsErrorEqual(t *testing.T) {
	tests := []struct {
		name string
		a    error
		b    error
		want bool
	}{
		{
			name: "Both nil",
			a:    nil,
			b:    nil,
			want: true,
		},
		{
			name: "Error a nil",
			a:    nil,
			b:    fmt.Errorf("Error"),
			want: false,
		},
		{
			name: "Error b Nil",
			a:    fmt.Errorf("Error"),
			b:    nil,
			want: false,
		},
		{
			name: "Both different error",
			a:    fmt.Errorf("Error 1"),
			b:    fmt.Errorf("Error 2"),
			want: false,
		},
		{
			name: "Same error",
			a:    fmt.Errorf("Error"),
			b:    fmt.Errorf("Error"),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsErrorEqual(tt.a, tt.b); got != tt.want {
				t.Errorf("IsErrorDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
