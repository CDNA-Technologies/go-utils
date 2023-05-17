package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ExchangeRateAPIBaseURL = "https://api.exchangerate.host"
	LatestRateEndpoint     = "/latest"
)

type LatestRatesRequest struct {
	BaseCurrency   string
	WantCurrencies []string
	Amount         float64
	RoundPlaces    int32
}

type LatestRatesResponse struct {
	BaseCurrency string             `json:"base,omitempty"`
	Date         string             `json:"date,omitempty"`
	Rates        map[string]float32 `json:"rates,omitempty"`
}

func latestRateUrl(req LatestRatesRequest) string {
	url := ExchangeRateAPIBaseURL + LatestRateEndpoint

	url += fmt.Sprintf("?amount=%f", req.Amount)
	url += fmt.Sprintf("&places=%d", req.RoundPlaces)

	if req.BaseCurrency != "" {
		url += fmt.Sprintf("&base=%s", req.BaseCurrency)
	}

	if req.WantCurrencies != nil && len(req.WantCurrencies) > 0 {
		url += "&symbols="
		for _, c := range req.WantCurrencies {
			url += c + ","
		}
	}

	// Remove the extra ',' at the end
	return url[0 : len(url)-1]
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

	var res LatestRatesResponse
	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		return LatestRatesResponse{}, fmt.Errorf("unable to parse response body: %v, error: %v", res, res)
	}

	return res, nil
}
