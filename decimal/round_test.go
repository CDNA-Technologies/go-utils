package decimal

import (
	"errors"
	"fmt"
	"testing"
)

func TestRound(t *testing.T) {
	tests := []struct {
		val     interface{}
		places  int32
		want    float64
		wantErr error
	}{
		// Test cases for float32 input
		{val: float32(1.234), places: 2, want: 1.23, wantErr: nil},
		{val: float32(3.14159), places: 2, want: 3.14, wantErr: nil},
		{val: float32(3.1459), places: 1, want: 3.1, wantErr: nil},
		{val: float32(-3.1459), places: 2, want: -3.15, wantErr: nil},
		{val: float32(3.1459), places: -1, want: 3, wantErr: nil},

		// Test cases for float64 input
		{val: float64(1.234), places: 2, want: 1.23, wantErr: nil},
		{val: 3.14159, places: 2, want: 3.14, wantErr: nil},
		{val: 3.1459, places: 1, want: 3.1, wantErr: nil},
		{val: -3.1459, places: 2, want: -3.15, wantErr: nil},
		{val: 3.1459, places: -1, want: 3, wantErr: nil},

		// Test cases for string input
		{val: "3.14159", places: 2, want: 3.14, wantErr: nil},
		{val: "3.1459", places: 1, want: 3.1, wantErr: nil},
		{val: "-3.1459", places: 2, want: -3.15, wantErr: nil},
		{val: "3.1459", places: -1, want: 3, wantErr: nil},
		{val: "1.234", places: 2, want: 1.23, wantErr: nil},
		{val: "1.234", places: -2, want: 1, wantErr: nil},
		{val: "invalid", places: 2, want: 0, wantErr: errors.New("strconv.ParseFloat: parsing \"invalid\": invalid syntax")},
		{val: "not a float", places: 2, wantErr: errors.New("strconv.ParseFloat: parsing \"not a float\": invalid syntax")},

		// Nil input
		{val: nil, places: 2, want: 0, wantErr: errors.New("value is not a float or string")},
	}

	for _, input := range tests {
		t.Run(fmt.Sprintf("TestRound(%v, %d)", input.val, input.places), func(t *testing.T) {
			got, err := Round(input.val, input.places)

			if err != nil {
				if input.wantErr == nil || input.wantErr.Error() != err.Error() {
					t.Errorf("Round(%v, %d) got error = %v, wantErr %v", input.val, input.places, err, input.wantErr)
				}
			}

			if input.want != got {
				t.Errorf("Round(%v, %d) = %v, want %v", input.val, input.places, got, input.want)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		val     interface{}
		places  int32
		want    float64
		wantErr error
	}{
		// Test cases for float32 input
		{val: float32(1.234), places: 2, want: 1.24, wantErr: nil},
		{val: float32(3.14159), places: 2, want: 3.15, wantErr: nil},
		{val: float32(3.1459), places: 1, want: 3.2, wantErr: nil},
		{val: float32(-3.1459), places: 2, want: -3.14, wantErr: nil},
		{val: float32(3.1459), places: -1, want: 4, wantErr: nil},

		// Test cases for float64 input
		{val: float64(1.234), places: 2, want: 1.24, wantErr: nil},
		{val: 3.14159, places: 2, want: 3.15, wantErr: nil},
		{val: 3.1459, places: 1, want: 3.2, wantErr: nil},
		{val: -3.1459, places: 2, want: -3.14, wantErr: nil},
		{val: 3.1459, places: -1, want: 4, wantErr: nil},

		// Test cases for string input
		{val: "3.14159", places: 2, want: 3.15, wantErr: nil},
		{val: "3.1459", places: 1, want: 3.2, wantErr: nil},
		{val: "-3.1459", places: 2, want: -3.14, wantErr: nil},
		{val: "3.1459", places: -1, want: 4, wantErr: nil},
		{val: "1.234", places: 2, want: 1.24, wantErr: nil},
		{val: "1.234", places: -2, want: 2, wantErr: nil},
		{val: "invalid", places: 2, want: 0, wantErr: errors.New("strconv.ParseFloat: parsing \"invalid\": invalid syntax")},
		{val: "not a float", places: 2, wantErr: errors.New("strconv.ParseFloat: parsing \"not a float\": invalid syntax")},

		// Nil input
		{val: nil, places: 2, want: 0, wantErr: errors.New("value is not a float or string")},
	}

	for _, input := range tests {
		t.Run(fmt.Sprintf("TestCeil(%v, %d)", input.val, input.places), func(t *testing.T) {
			got, err := Ceil(input.val, input.places)

			if err != nil {
				if input.wantErr == nil || input.wantErr.Error() != err.Error() {
					t.Errorf("Ceil(%v, %d) got error = %v, wantErr %v", input.val, input.places, err, input.wantErr)
				}
			}

			if input.want != got {
				t.Errorf("Ceil(%v, %d) = %v, want %v", input.val, input.places, got, input.want)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		val     interface{}
		places  int32
		want    float64
		wantErr error
	}{
		// Test cases for float32 input
		{val: float32(1.234), places: 2, want: 1.23, wantErr: nil},
		{val: float32(3.14159), places: 2, want: 3.14, wantErr: nil},
		{val: float32(3.1459), places: 1, want: 3.1, wantErr: nil},
		{val: float32(-3.1459), places: 2, want: -3.15, wantErr: nil},
		{val: float32(3.1459), places: -1, want: 3, wantErr: nil},

		// Test cases for float64 input
		{val: float64(1.234), places: 2, want: 1.23, wantErr: nil},
		{val: 3.14159, places: 2, want: 3.14, wantErr: nil},
		{val: 3.1459, places: 1, want: 3.1, wantErr: nil},
		{val: -3.1459, places: 2, want: -3.15, wantErr: nil},
		{val: 3.1459, places: -1, want: 3, wantErr: nil},

		// Test cases for string input
		{val: "3.14159", places: 2, want: 3.14, wantErr: nil},
		{val: "3.1459", places: 1, want: 3.1, wantErr: nil},
		{val: "-3.1459", places: 2, want: -3.15, wantErr: nil},
		{val: "3.1459", places: -1, want: 3, wantErr: nil},
		{val: "1.234", places: 2, want: 1.23, wantErr: nil},
		{val: "1.234", places: -2, want: 1, wantErr: nil},
		{val: "invalid", places: 2, want: 0, wantErr: errors.New("strconv.ParseFloat: parsing \"invalid\": invalid syntax")},
		{val: "not a float", places: 2, wantErr: errors.New("strconv.ParseFloat: parsing \"not a float\": invalid syntax")},

		// Nil input
		{val: nil, places: 2, want: 0, wantErr: errors.New("value is not a float or string")},
	}

	for _, input := range tests {
		t.Run(fmt.Sprintf("TestFloor(%v, %d)", input.val, input.places), func(t *testing.T) {
			got, err := Floor(input.val, input.places)

			if err != nil {
				if input.wantErr == nil || input.wantErr.Error() != err.Error() {
					t.Errorf("Floor(%v, %d) got error = %v, wantErr %v", input.val, input.places, err, input.wantErr)
				}
			}

			if input.want != got {
				t.Errorf("Floor(%v, %d) = %v, want %v", input.val, input.places, got, input.want)
			}
		})
	}
}
