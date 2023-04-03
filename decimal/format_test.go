package decimal

import (
	"errors"
	"fmt"
	"testing"
)

func TestRoundAndFormat(t *testing.T) {
	tests := []struct {
		val     interface{}
		places  int32
		want    string
		wantErr error
	}{
		// Basic test cases
		{1.234, 2, "1.23", nil},
		{1.235, 2, "1.24", nil},
		{1.236, 2, "1.24", nil},
		{1.225, 2, "1.23", nil},
		{1.215, 2, "1.22", nil},
		{1.225, 1, "1.2", nil},
		{1.234, 0, "1", nil},
		{1234, 2, "1234.00", nil},
		{float32(1.234), 2, "1.23", nil},
		{"1.234", 2, "1.23", nil},
		{"1", 2, "1.00", nil},
		{"1.2345", 2, "1.23", nil},
		{"-1.234", 2, "-1.23", nil},
		{"1.239", 0, "1", nil},
		{"-1.239", 0, "-1", nil},
		{0, 2, "0.00", nil},

		// Test cases with padding/trimming zeros
		{1.2, 4, "1.2000", nil},
		{1.234, 4, "1.2340", nil},
		{1.200, 4, "1.2000", nil},
		{1.000, 4, "1.0000", nil},
		{1, 4, "1.0000", nil},
		{"1.2", 4, "1.2000", nil},
		{"1.234", 4, "1.2340", nil},
		{"1.200", 4, "1.2000", nil},
		{"1.000", 4, "1.0000", nil},
		{"1", 4, "1.0000", nil},
		{"-1.2", 4, "-1.2000", nil},
		{"-1.234", 4, "-1.2340", nil},
		{"-1.200", 4, "-1.2000", nil},
		{"-1.000", 4, "-1.0000", nil},
		{"-1", 4, "-1.0000", nil},
		{"0", 4, "0.0000", nil},

		// Error test cases
		{"invalid", 2, "", errors.New("strconv.ParseFloat: parsing \"invalid\": invalid syntax")},
		{complex(1, 2), 2, "", errors.New("value is not a float or string")},
	}

	for _, input := range tests {
		t.Run(fmt.Sprintf("TestRoundAndFormat(%v, %d)", input.val, input.places), func(t *testing.T) {
			got, err := RoundAndFormat(input.val, input.places)

			if err != nil {
				if input.wantErr == nil || input.wantErr.Error() != err.Error() {
					t.Errorf("RoundAndFormat(%v, %d) got error = %v, wantErr %v", input.val, input.places, err, input.wantErr)
				}
			}

			if input.want != got {
				t.Errorf("RoundAndFormat(%v, %d) = %v, want %v", input.val, input.places, got, input.want)
			}
		})
	}
}

func TestCeilAndFormat(t *testing.T) {
	tests := []struct {
		val     interface{}
		places  int32
		want    string
		wantErr error
	}{
		// Basic test cases
		{1.234, 2, "1.24", nil},
		{1.235, 2, "1.24", nil},
		{1.236, 2, "1.24", nil},
		{1.225, 2, "1.23", nil},
		{1.215, 2, "1.22", nil},
		{1.225, 1, "1.3", nil},
		{1.234, 0, "2", nil},
		{1234, 2, "1234.00", nil},
		{float32(1.234), 2, "1.24", nil},
		{"1.234", 2, "1.24", nil},
		{"1", 2, "1.00", nil},
		{"1.2345", 2, "1.24", nil},
		{"-1.234", 2, "-1.23", nil},
		{"1.239", 0, "2", nil},
		{"-1.239", 0, "-1", nil},
		{0, 2, "0.00", nil},

		// Test cases with padding/trimming zeros
		{1.2, 4, "1.2000", nil},
		{1.234, 4, "1.2340", nil},
		{1.200, 4, "1.2000", nil},
		{1.000, 4, "1.0000", nil},
		{1, 4, "1.0000", nil},
		{"1.2", 4, "1.2000", nil},
		{"1.234", 4, "1.2340", nil},
		{"1.200", 4, "1.2000", nil},
		{"1.000", 4, "1.0000", nil},
		{"1", 4, "1.0000", nil},
		{"-1.2", 4, "-1.2000", nil},
		{"-1.234", 4, "-1.2340", nil},
		{"-1.200", 4, "-1.2000", nil},
		{"-1.000", 4, "-1.0000", nil},
		{"-1", 4, "-1.0000", nil},
		{"0", 4, "0.0000", nil},

		// Error test cases
		{"invalid", 2, "", errors.New("strconv.ParseFloat: parsing \"invalid\": invalid syntax")},
		{complex(1, 2), 2, "", errors.New("value is not a float or string")},
	}

	for _, input := range tests {
		t.Run(fmt.Sprintf("TestCeilAndFormat(%v, %d)", input.val, input.places), func(t *testing.T) {
			got, err := CeilAndFormat(input.val, input.places)

			if err != nil {
				if input.wantErr == nil || input.wantErr.Error() != err.Error() {
					t.Errorf("CeilAndFormat(%v, %d) got error = %v, wantErr %v", input.val, input.places, err, input.wantErr)
				}
			}

			if input.want != got {
				t.Errorf("CeilAndFormat(%v, %d) = %v, want %v", input.val, input.places, got, input.want)
			}
		})
	}
}

func TestFloorAndFormat(t *testing.T) {
	tests := []struct {
		val     interface{}
		places  int32
		want    string
		wantErr error
	}{
		// Basic test cases
		{1.234, 2, "1.23", nil},
		{1.235, 2, "1.23", nil},
		{1.236, 2, "1.23", nil},
		{1.225, 2, "1.22", nil},
		{1.215, 2, "1.21", nil},
		{1.225, 1, "1.2", nil},
		{1.234, 0, "1", nil},
		{1234, 2, "1234.00", nil},
		{float32(1.234), 2, "1.23", nil},
		{"1.234", 2, "1.23", nil},
		{"1", 2, "1.00", nil},
		{"1.2345", 2, "1.23", nil},
		{"-1.234", 2, "-1.24", nil},
		{"1.239", 0, "1", nil},
		{"-1.239", 0, "-2", nil},
		{0, 2, "0.00", nil},

		// Test cases with padding/trimming zeros
		{1.2, 4, "1.2000", nil},
		{1.234, 4, "1.2340", nil},
		{1.200, 4, "1.2000", nil},
		{1.000, 4, "1.0000", nil},
		{1, 4, "1.0000", nil},
		{"1.2", 4, "1.2000", nil},
		{"1.234", 4, "1.2340", nil},
		{"1.200", 4, "1.2000", nil},
		{"1.000", 4, "1.0000", nil},
		{"1", 4, "1.0000", nil},
		{"-1.2", 4, "-1.2000", nil},
		{"-1.234", 4, "-1.2340", nil},
		{"-1.200", 4, "-1.2000", nil},
		{"-1.000", 4, "-1.0000", nil},
		{"-1", 4, "-1.0000", nil},
		{"0", 4, "0.0000", nil},

		// Error test cases
		{"invalid", 2, "", errors.New("strconv.ParseFloat: parsing \"invalid\": invalid syntax")},
		{complex(1, 2), 2, "", errors.New("value is not a float or string")},
	}

	for _, input := range tests {
		t.Run(fmt.Sprintf("TestFloorAndFormat(%v, %d)", input.val, input.places), func(t *testing.T) {
			got, err := FloorAndFormat(input.val, input.places)

			if err != nil {
				if input.wantErr == nil || input.wantErr.Error() != err.Error() {
					t.Errorf("FloorAndFormat(%v, %d) got error = %v, wantErr %v", input.val, input.places, err, input.wantErr)
				}
			}

			if input.want != got {
				t.Errorf("FloorAndFormat(%v, %d) = %v, want %v", input.val, input.places, got, input.want)
			}
		})
	}
}
