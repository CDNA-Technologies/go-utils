package decimal

import (
	"strconv"
)

func RoundAndFormat(val interface{}, places int32) (string, error) {
	rounded, err := Round(val, places)
	if err != nil {
		return "", err
	}
	return strconv.FormatFloat(rounded, 'f', int(places), 64), nil
}

func CeilAndFormat(val interface{}, places int32) (string, error) {
	rounded, err := Ceil(val, places)
	if err != nil {
		return "", err
	}
	return strconv.FormatFloat(rounded, 'f', int(places), 64), nil
}

func FloorAndFormat(val interface{}, places int32) (string, error) {
	rounded, err := Floor(val, places)
	if err != nil {
		return "", err
	}
	return strconv.FormatFloat(rounded, 'f', int(places), 64), nil
}
