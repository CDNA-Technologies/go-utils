package decimal

import (
	"errors"
	"math"
	"strconv"
)

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
