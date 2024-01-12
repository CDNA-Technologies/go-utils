package decimal

type currencyConversion struct {
	CountryCode          string
	Language             string
	IsConversionRequired bool
}

var countryList = []currencyConversion{
	{
		CountryCode:          "ID",
		Language:             "id",
		IsConversionRequired: true,
	},
}
