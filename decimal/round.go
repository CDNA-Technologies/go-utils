package decimal

import (
	"errors"
	"math"
	"strconv"
)

/**
	Rounds the given value to the given number of decimal places.

	Example: Round(10.576,2) = 10.58, Round(6.7816, 3) = 6.781
**/
func Round(val interface{}, places int32) (float64, error) {
	switch v := val.(type) {
	case int:
		fval := float64(v)
		return round(fval, places), nil
	case int32:
		fval := float64(v)
		return round(fval, places), nil
	case int64:
		fval := float64(v)
		return round(fval, places), nil
	case float32:
		fval := float64(v)
		return round(fval, places), nil
	case float64:
		return round(v, places), nil
	case string:
		fval, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
		return round(fval, places), nil
	default:
		return 0, errors.New("value is not a float or string")
	}
}

func round(val float64, places int32) float64 {
	if places < 0 {
		places = 0
	}
	shift := math.Pow(10, float64(places))
	return math.Round(val*shift) / shift
}

/**
	Calculates the smallest integer value greater than or equal to the input value, rounded up to the given number of decimal places.

	Example: Ceil(10.526,2) = 10.53, Round(6.7816, 3) = 6.782
**/
func Ceil(val interface{}, places int32) (float64, error) {
	switch v := val.(type) {
	case int:
		fval := float64(v)
		return ceil(fval, places), nil
	case int32:
		fval := float64(v)
		return ceil(fval, places), nil
	case int64:
		fval := float64(v)
		return ceil(fval, places), nil
	case float32:
		fval := float64(v)
		return ceil(fval, places), nil
	case float64:
		return ceil(v, places), nil
	case string:
		fval, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
		return ceil(fval, places), nil
	default:
		return 0, errors.New("value is not a float or string")
	}
}

func ceil(val float64, places int32) float64 {
	if places < 0 {
		places = 0
	}
	shift := math.Pow(10, float64(places))
	return math.Ceil(val*shift) / shift
}

/**
	Calculates the smallest integer value lesser than or equal to the input value, rounded up to the given number of decimal places.

	Example: Ceil(10.526,2) = 10.52, Round(6.7816, 3) = 6.781
**/
func Floor(val interface{}, places int32) (float64, error) {
	switch v := val.(type) {
	case int:
		fval := float64(v)
		return floor(fval, places), nil
	case int32:
		fval := float64(v)
		return floor(fval, places), nil
	case int64:
		fval := float64(v)
		return floor(fval, places), nil
	case float32:
		fval := float64(v)
		return floor(fval, places), nil
	case float64:
		return floor(v, places), nil
	case string:
		fval, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
		return floor(fval, places), nil
	default:
		return 0, errors.New("value is not a float or string")
	}
}

func floor(val float64, places int32) float64 {
	if places < 0 {
		places = 0
	}
	shift := math.Pow(10, float64(places))
	return math.Floor(val*shift) / shift
}
