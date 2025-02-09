package pricefeed

import (
	"encoding/json"
	"log"
	"net/http"
)

type PriceData struct {
	Crypto string  `json:"crypto"`
	Price  float64 `json:"price"`
}

func FetchPrice(crypto string) (float64, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + crypto + "&vs_currencies=usd"

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to fetch price: %v", err)
		return 0, err
	}
	defer resp.Body.Close()

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Failed to decode price data: %v", err)
		return 0, err
	}

	price, ok := result[crypto]["usd"]
	if !ok {
		return 0, nil
	}

	return price, nil
}