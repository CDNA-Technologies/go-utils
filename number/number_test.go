package number

import (
	"fmt"
	"math"
	"testing"
)

func TestIsWhole(t *testing.T) {
	tests := []struct {
		input    float64
		expected bool
	}{
		{
			input:    0.000,
			expected: true,
		},
		{
			input:    2.90,
			expected: false,
		},
		{
			input:    1.9,
			expected: false,
		},
		{
			input:    1,
			expected: true,
		},
		{
			input:    0.11,
			expected: false,
		},
		{
			input:    -0.11,
			expected: false,
		},
		{
			input:    -1,
			expected: true,
		},
		{
			input:    math.Inf(-1),
			expected: false,
		},
		{
			input:    math.Inf(1),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("test isWholeNumber for %f", test.input), func(t *testing.T) {
			resp := IsWhole(test.input)

			if resp != test.expected {
				t.Errorf("unexpected result %v", resp)
			}
		})
	}
}
