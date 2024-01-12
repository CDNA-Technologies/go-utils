package decimal

import (
	"errors"
	"fmt"
	"testing"
)

func TestRoundAndFormat(t *testing.T) {
	tests := []struct {
		val         interface{}
		places      int32
		countryCode string
		want        string
		wantErr     error
	}{
		// Basic test cases
		{1.234, 2, "IN", "1.23", nil},
		{1.235, 2, "IN", "1.24", nil},
		{1.236, 2, "IN", "1.24", nil},
		{1.225, 2, "IN", "1.23", nil},
		{1.215, 2, "IN", "1.22", nil},
		{1.225, 1, "IN", "1.2", nil},
		{1.234, 0, "IN", "1", nil},
		{1234, 2, "IN", "1234.00", nil},
		{float32(1.234), 2, "IN", "1.23", nil},
		{"1.234", 2, "IN", "1.23", nil},
		{"1", 2, "IN", "1.00", nil},
		{"1.2345", 2, "IN", "1.23", nil},
		{"-1.234", 2, "IN", "-1.23", nil},
		{"1.239", 0, "IN", "1", nil},
		{"-1.239", 0, "IN", "-1", nil},
		{0, 2, "IN", "0.00", nil},

		// Test cases with padding/trimming zeros
		{1.2, 4, "IN", "1.2000", nil},
		{1.234, 4, "IN", "1.2340", nil},
		{1.200, 4, "IN", "1.2000", nil},
		{1.000, 4, "IN", "1.0000", nil},
		{1, 4, "IN", "1.0000", nil},
		{"1.2", 4, "IN", "1.2000", nil},
		{"1.234", 4, "IN", "1.2340", nil},
		{"1.200", 4, "IN", "1.2000", nil},
		{"1.000", 4, "IN", "1.0000", nil},
		{"1", 4, "IN", "1.0000", nil},
		{"-1.2", 4, "IN", "-1.2000", nil},
		{"-1.234", 4, "IN", "-1.2340", nil},
		{"-1.200", 4, "IN", "-1.2000", nil},
		{"-1.000", 4, "IN", "-1.0000", nil},
		{"-1", 4, "IN", "-1.0000", nil},
		{"0", 4, "IN", "0.0000", nil},
		{"0", -4, "IN", "0", nil},

		// Error test cases
		{"invalid", 2, "IN", "", errors.New("strconv.ParseFloat: parsing \"invalid\": invalid syntax")},
		{complex(1, 2), 2, "IN", "", errors.New("value is not a float or string")},

		// Basic test cases for indonesia
		{1.234, 2, "ID", "1,23", nil},
		{1.235, 2, "ID", "1,24", nil},
		{1.236, 2, "ID", "1,24", nil},
		{1.225, 2, "ID", "1,23", nil},
		{1.215, 2, "ID", "1,22", nil},
		{1.225, 1, "ID", "1,2", nil},
		{1.234, 0, "ID", "1", nil},
		{1234, 2, "ID", "1.234,00", nil},
		{float32(1.234), 2, "ID", "1,23", nil},
		{"1.234", 2, "ID", "1,23", nil},
		{"1", 2, "ID", "1,00", nil},
		{"1.2345", 2, "ID", "1,23", nil},
		{"-1.234", 2, "ID", "-1,23", nil},
		{"1.239", 0, "ID", "1", nil},
		{"-1.239", 0, "ID", "-1", nil},
		{0, 2, "ID", "0,00", nil},

		// Test cases with padding/trimming zeros indonesia
		{1.2, 4, "ID", "1,2000", nil},
		{1.234, 4, "ID", "1,2340", nil},
		{1.200, 4, "ID", "1,2000", nil},
		{1.000, 4, "ID", "1,0000", nil},
		{1, 4, "ID", "1,0000", nil},
		{"1.2", 4, "ID", "1,2000", nil},
		{"1.234", 4, "ID", "1,2340", nil},
		{"1.200", 4, "ID", "1,2000", nil},
		{"1.000", 4, "ID", "1,0000", nil},
		{"1", 4, "ID", "1,0000", nil},
		{"-1.2", 4, "ID", "-1,2000", nil},
		{"-1.234", 4, "ID", "-1,2340", nil},
		{"-1.200", 4, "ID", "-1,2000", nil},
		{"-1.000", 4, "ID", "-1,0000", nil},
		{"-1", 4, "ID", "-1,0000", nil},
		{"0", 4, "ID", "0,0000", nil},
		{"0", -4, "ID", "0", nil},

		// Error test cases
		{"invalid", 2, "ID", "", errors.New("strconv.ParseFloat: parsing \"invalid\": invalid syntax")},
		{complex(1, 2), 2, "ID", "", errors.New("value is not a float or string")},
	}

	for _, input := range tests {
		t.Run(fmt.Sprintf("TestRoundAndFormat(%v, %d)", input.val, input.places), func(t *testing.T) {
			got, err := RoundAndFormat(input.val, input.places, input.countryCode)

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
		val         interface{}
		places      int32
		countryCode string
		want        string
		wantErr     error
	}{
		// Basic test cases
		{1.234, 2, "IN", "1.24", nil},
		{1.235, 2, "IN", "1.24", nil},
		{1.236, 2, "IN", "1.24", nil},
		{1.225, 2, "IN", "1.23", nil},
		{1.215, 2, "IN", "1.22", nil},
		{1.225, 1, "IN", "1.3", nil},
		{1.234, 0, "IN", "2", nil},
		{1234, 2, "IN", "1234.00", nil},
		{float32(1.234), 2, "IN", "1.24", nil},
		{"1.234", 2, "IN", "1.24", nil},
		{"1", 2, "IN", "1.00", nil},
		{"1.2345", 2, "IN", "1.24", nil},
		{"-1.234", 2, "IN", "-1.23", nil},
		{"1.239", 0, "IN", "2", nil},
		{"-1.239", 0, "IN", "-1", nil},
		{0, 2, "IN", "0.00", nil},

		// Test cases with padding/trimming zeros
		{1.2, 4, "IN", "1.2000", nil},
		{1.234, 4, "IN", "1.2340", nil},
		{1.200, 4, "IN", "1.2000", nil},
		{1.000, 4, "IN", "1.0000", nil},
		{1, 4, "IN", "1.0000", nil},
		{"1.2", 4, "IN", "1.2000", nil},
		{"1.234", 4, "IN", "1.2340", nil},
		{"1.200", 4, "IN", "1.2000", nil},
		{"1.000", 4, "IN", "1.0000", nil},
		{"1", 4, "IN", "1.0000", nil},
		{"-1.2", 4, "IN", "-1.2000", nil},
		{"-1.234", 4, "IN", "-1.2340", nil},
		{"-1.200", 4, "IN", "-1.2000", nil},
		{"-1.000", 4, "IN", "-1.0000", nil},
		{"-1", 4, "IN", "-1.0000", nil},
		{"0", 4, "IN", "0.0000", nil},
		{"1", -4, "IN", "1", nil},

		// Error test cases
		{"invalid", 2, "IN", "", errors.New("strconv.ParseFloat: parsing \"invalid\": invalid syntax")},
		{complex(1, 2), 2, "IN", "", errors.New("value is not a float or string")},

		// Basic test cases indonesia
		{1.234, 2, "ID", "1,24", nil},
		{1.235, 2, "ID", "1,24", nil},
		{1.236, 2, "ID", "1,24", nil},
		{1.225, 2, "ID", "1,23", nil},
		{1.215, 2, "ID", "1,22", nil},
		{1.225, 1, "ID", "1,3", nil},
		{1.234, 0, "ID", "2", nil},
		{1234, 2, "ID", "1.234,00", nil},
		{float32(1.234), 2, "ID", "1,24", nil},
		{"1.234", 2, "ID", "1,24", nil},
		{"1", 2, "ID", "1,00", nil},
		{"1.2345", 2, "ID", "1,24", nil},
		{"-1.234", 2, "ID", "-1,23", nil},
		{"1.239", 0, "ID", "2", nil},
		{"-1.239", 0, "ID", "-1", nil},
		{0, 2, "ID", "0,00", nil},

		// Test cases with padding/trimming zeros indonesia
		{1.2, 4, "ID", "1,2000", nil},
		{1.234, 4, "ID", "1,2340", nil},
		{1.200, 4, "ID", "1,2000", nil},
		{1.000, 4, "ID", "1,0000", nil},
		{1, 4, "ID", "1,0000", nil},
		{"1.2", 4, "ID", "1,2000", nil},
		{"1.234", 4, "ID", "1,2340", nil},
		{"1.200", 4, "ID", "1,2000", nil},
		{"1.000", 4, "ID", "1,0000", nil},
		{"1", 4, "ID", "1,0000", nil},
		{"-1.2", 4, "ID", "-1,2000", nil},
		{"-1.234", 4, "ID", "-1,2340", nil},
		{"-1.200", 4, "ID", "-1,2000", nil},
		{"-1.000", 4, "ID", "-1,0000", nil},
		{"-1", 4, "ID", "-1,0000", nil},
		{"0", 4, "ID", "0,0000", nil},
		{"1", -4, "ID", "1", nil},

		// Error test cases
		{"invalid", 2, "IN", "", errors.New("strconv.ParseFloat: parsing \"invalid\": invalid syntax")},
		{complex(1, 2), 2, "IN", "", errors.New("value is not a float or string")},
	}

	for _, input := range tests {
		t.Run(fmt.Sprintf("TestCeilAndFormat(%v, %d)", input.val, input.places), func(t *testing.T) {
			got, err := CeilAndFormat(input.val, input.places, input.countryCode)

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
		val         interface{}
		places      int32
		countryCode string
		want        string
		wantErr     error
	}{
		// Basic test cases
		{1.234, 2, "IN", "1.23", nil},
		{1.235, 2, "IN", "1.23", nil},
		{1.236, 2, "IN", "1.23", nil},
		{1.225, 2, "IN", "1.22", nil},
		{1.215, 2, "IN", "1.21", nil},
		{1.225, 1, "IN", "1.2", nil},
		{1.234, 0, "IN", "1", nil},
		{1234, 2, "IN", "1234.00", nil},
		{float32(1.234), 2, "IN", "1.23", nil},
		{"1.234", 2, "IN", "1.23", nil},
		{"1", 2, "IN", "1.00", nil},
		{"1.2345", 2, "IN", "1.23", nil},
		{"-1.234", 2, "IN", "-1.24", nil},
		{"1.239", 0, "IN", "1", nil},
		{"-1.239", 0, "IN", "-2", nil},
		{0, 2, "IN", "0.00", nil},

		// Test cases with padding/trimming zeros
		{1.2, 4, "IN", "1.2000", nil},
		{1.234, 4, "IN", "1.2340", nil},
		{1.200, 4, "IN", "1.2000", nil},
		{1.000, 4, "IN", "1.0000", nil},
		{1, 4, "IN", "1.0000", nil},
		{"1.2", 4, "IN", "1.2000", nil},
		{"1.234", 4, "IN", "1.2340", nil},
		{"1.200", 4, "IN", "1.2000", nil},
		{"1.000", 4, "IN", "1.0000", nil},
		{"1", 4, "IN", "1.0000", nil},
		{"-1.2", 4, "IN", "-1.2000", nil},
		{"-1.234", 4, "IN", "-1.2340", nil},
		{"-1.200", 4, "IN", "-1.2000", nil},
		{"-1.000", 4, "IN", "-1.0000", nil},
		{"-1", 4, "IN", "-1.0000", nil},
		{"0", 4, "IN", "0.0000", nil},
		{"0", -4, "IN", "0", nil},

		// Error test cases
		{"invalid", 2, "IN", "", errors.New("strconv.ParseFloat: parsing \"invalid\": invalid syntax")},
		{complex(1, 2), 2, "IN", "", errors.New("value is not a float or string")},

		// Basic test cases Indonesia
		{1.234, 2, "ID", "1,23", nil},
		{1.235, 2, "ID", "1,23", nil},
		{1.236, 2, "ID", "1,23", nil},
		{1.225, 2, "ID", "1,22", nil},
		{1.215, 2, "ID", "1,21", nil},
		{1.225, 1, "ID", "1,2", nil},
		{1.234, 0, "ID", "1", nil},
		{1234, 2, "ID", "1.234,00", nil},
		{float32(1.234), 2, "ID", "1,23", nil},
		{"1.234", 2, "ID", "1,23", nil},
		{"1", 2, "ID", "1,00", nil},
		{"1.2345", 2, "ID", "1,23", nil},
		{"-1.234", 2, "ID", "-1,24", nil},
		{"1.239", 0, "ID", "1", nil},
		{"-1.239", 0, "ID", "-2", nil},
		{0, 2, "ID", "0,00", nil},

		// Test cases with padding/trimming zeros Indonesia
		{1.2, 4, "ID", "1,2000", nil},
		{1.234, 4, "ID", "1,2340", nil},
		{1.200, 4, "ID", "1,2000", nil},
		{1.000, 4, "ID", "1,0000", nil},
		{1, 4, "ID", "1,0000", nil},
		{"1.2", 4, "ID", "1,2000", nil},
		{"1.234", 4, "ID", "1,2340", nil},
		{"1.200", 4, "ID", "1,2000", nil},
		{"1.000", 4, "ID", "1,0000", nil},
		{"1", 4, "ID", "1,0000", nil},
		{"-1.2", 4, "ID", "-1,2000", nil},
		{"-1.234", 4, "ID", "-1,2340", nil},
		{"-1.200", 4, "ID", "-1,2000", nil},
		{"-1.000", 4, "ID", "-1,0000", nil},
		{"-1", 4, "ID", "-1,0000", nil},
		{"0", 4, "ID", "0,0000", nil},
		{"0", -4, "ID", "0", nil},

		// Error test cases
		{"invalid", 2, "ID", "", errors.New("strconv.ParseFloat: parsing \"invalid\": invalid syntax")},
		{complex(1, 2), 2, "ID", "", errors.New("value is not a float or string")},
	}

	for _, input := range tests {
		t.Run(fmt.Sprintf("TestFloorAndFormat(%v, %d)", input.val, input.places), func(t *testing.T) {
			got, err := FloorAndFormat(input.val, input.places, input.countryCode)

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

func TestFormatCountryWise(t *testing.T) {
	tests := []struct {
		val         interface{}
		countryCode string
		want        string
		wantErr     error
	}{
		// Basic test cases decimal places doesn't matter use it only for whole numbers
		{1234, "IN", "1234", nil},
		{1234.00, "IN", "1234", nil},

		// Basic test cases for indonesia
		{1234, "ID", "1.234", nil},
		// Test cases with padding/trimming zeros indonesia
		{1111111.2, "ID", "1.111.111", nil},
		{1111111.5, "ID", "1.111.112", nil},
	}

	for _, input := range tests {
		t.Run(fmt.Sprintf("TestRoundAndFormat(%v)", input.val), func(t *testing.T) {
			got, err := FormatWholeNumberCountryWise(input.val, input.countryCode)

			if err != nil {
				if input.wantErr == nil || input.wantErr.Error() != err.Error() {
					t.Errorf("RoundAndFormat(%v) got error = %v, wantErr %v", input.val, err, input.wantErr)
				}
			}

			if input.want != got {
				t.Errorf("FormatWholeNumberCountryWise(%v) = %v, want %v", input.val, got, input.want)
			}
		})
	}
}
