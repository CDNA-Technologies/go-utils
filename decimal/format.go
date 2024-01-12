package decimal

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

/**
	Round the value to the given number of places and format the result to the same number of decimal places.

	Example: RoundAndFormat(45.2489, 2, IN) = 45.25,  RoundAndFormat(45.2, 3, IN) = 45.200
**/
func RoundAndFormat(val interface{}, places int32, countryCode string) (string, error) {
	if places < 0 {
		places = 0
	}
	rounded, err := Round(val, places)
	if err != nil {
		return "", err
	}
	for _, country := range countryList {
		if countryCode == country.CountryCode && country.IsConversionRequired {
			lang := language.MustParse(country.Language)
			dec := number.Decimal(rounded, number.Scale(int(places)))
			return message.NewPrinter(lang).Sprintf("%v", dec), nil
		}
	}
	return fmt.Sprintf("%0.*f", places, rounded), nil
}

/**
	Ceils the value to the given number of places and format the result to the same number of decimal places.

	Example: CeilAndFormat(45.2413, 2, IN) = 45.25,  CeilAndFormat(-1.239, 0, IN) = -1
**/
func CeilAndFormat(val interface{}, places int32, countryCode string) (string, error) {
	if places < 0 {
		places = 0
	}
	rounded, err := Ceil(val, places)
	if err != nil {
		return "", err
	}
	for _, country := range countryList {
		if countryCode == country.CountryCode && country.IsConversionRequired {
			lang := language.MustParse(country.Language)
			dec := number.Decimal(rounded, number.Scale(int(places)))
			return message.NewPrinter(lang).Sprintf("%v", dec), nil
		}
	}
	return fmt.Sprintf("%0.*f", places, rounded), nil
}

/**
	Floors the value to the given number of places and format the result to the same number of decimal places.

	Example: CeilAndFormat(45.2393, 2, IN) = 45.23,  CeilAndFormat(-1.239, 0, IN) = -2
**/
func FloorAndFormat(val interface{}, places int32, countryCode string) (string, error) {
	if places < 0 {
		places = 0
	}
	rounded, err := Floor(val, places)
	if err != nil {
		return "", err
	}
	for _, country := range countryList {
		if countryCode == country.CountryCode && country.IsConversionRequired {
			lang := language.MustParse(country.Language)
			dec := number.Decimal(rounded, number.Scale(int(places)))
			return message.NewPrinter(lang).Sprintf("%v", dec), nil
		}
	}
	return fmt.Sprintf("%0.*f", places, rounded), nil
}

/**
	Use it only for Whole number to format it
	Example: FormatCountryWise(45.2393, 2, IN) = 45.2393,  FormatCountryWise(1234, 1, IN) = 1234
**/
func FormatWholeNumberCountryWise(val interface{}, countryCode string) (string, error) {
	for _, country := range countryList {
		if countryCode == country.CountryCode && country.IsConversionRequired {
			lang := language.MustParse(country.Language)
			dec := number.Decimal(val, number.Scale(0))
			return message.NewPrinter(lang).Sprintf("%v", dec), nil
		}
	}
	return fmt.Sprintf("%v", val), nil
}
