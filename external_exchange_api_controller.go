package main

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

func getCurrentBTCToUAHRate() float64 {
	var response apiResponse
	client := resty.New()
	resp, err := client.R().
		SetHeader("X-CoinAPI-Key", "1840BB94-23AA-4434-B89F-BD0D74FEFB32").
		Get("https://rest.coinapi.io/v1/exchangerate/BTC/UAH")
	if err != nil {
		return 0
	}
	json.Unmarshal(resp.Body(), &response)
	return response.Rate
}
