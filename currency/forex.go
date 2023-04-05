package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	ExchangeRateAPIBaseURL = "https://api.exchangerate.host"
	LatestRateEndpoint     = "/latest"
)

type LatestRatesRequest struct {
	BaseCurrency   interface{}
	WantCurrencies []interface{}
	Amount         interface{}
	RoundPlaces    interface{}
}

type LatestRatesResponse struct {
	BaseCurrency string
	Date         time.Duration
	Rates        map[string]float32
}

func latestRateUrl(req LatestRatesRequest) string {
	url := ExchangeRateAPIBaseURL + LatestRateEndpoint

	if req.Amount != nil {
		url += "?amount=" + req.Amount.(string)
	}

	if req.RoundPlaces != nil {
		url += "?places=" + req.RoundPlaces.(string)
	}

	if req.BaseCurrency != nil {
		url += "?base=" + req.BaseCurrency.(string)
	}

	if req.WantCurrencies != nil && len(req.WantCurrencies) != 0 {
		url += "?places="
		for _, c := range req.WantCurrencies {
			url += c.(string) + ","
		}
	}

	return url
}

func GetLatestRates(req LatestRatesRequest) (LatestRatesResponse, error) {
	request, err := http.NewRequest(http.MethodGet, latestRateUrl(req), nil)
	if err != nil {
		return LatestRatesResponse{}, err
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return LatestRatesResponse{}, nil
}
