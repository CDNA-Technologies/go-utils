package decimal

import (
	"fmt"
)

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
