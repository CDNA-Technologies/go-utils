package decimal

import (
	"fmt"
)

/**
	Round the value to the given number of places and format the result to the same number of decimal places.

	Example: RoundAndFormat(45.2489, 2) = 45.25,  RoundAndFormat(45.2, 3) = 45.200
**/
func RoundAndFormat(val interface{}, places int32) (string, error) {
	if places < 0 {
		places = 0
	}
	rounded, err := Round(val, places)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%0.*f", places, rounded), nil
}

/**
	Ceils the value to the given number of places and format the result to the same number of decimal places.

	Example: CeilAndFormat(45.2413, 2) = 45.25,  CeilAndFormat(-1.239, 0) = -1
**/
func CeilAndFormat(val interface{}, places int32) (string, error) {
	if places < 0 {
		places = 0
	}
	rounded, err := Ceil(val, places)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%0.*f", places, rounded), nil
}

/**
	Floors the value to the given number of places and format the result to the same number of decimal places.

	Example: CeilAndFormat(45.2393, 2) = 45.23,  CeilAndFormat(-1.239, 0) = -2
**/
func FloorAndFormat(val interface{}, places int32) (string, error) {
	if places < 0 {
		places = 0
	}
	rounded, err := Floor(val, places)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%0.*f", places, rounded), nil
}
