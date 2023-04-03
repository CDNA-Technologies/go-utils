package decimal

import (
	"fmt"
)

func RoundAndFormat(val interface{}, places int32) (string, error) {
	rounded, err := Round(val, places)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(fmt.Sprintf("%%.%df", places), rounded), nil
}

func CeilAndFormat(val interface{}, places int32) (string, error) {
	rounded, err := Ceil(val, places)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(fmt.Sprintf("%%.%df", places), rounded), nil
}

func FloorAndFormat(val interface{}, places int32) (string, error) {
	rounded, err := Floor(val, places)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(fmt.Sprintf("%%.%df", places), rounded), nil
}
