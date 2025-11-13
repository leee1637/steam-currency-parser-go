package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type MarketPrice struct {
	Success     bool   `json:"success"`
	LowestPrice string `json:"lowest_price"`
	Volume      string `json:"volume"`
	MedianPrice string `json:"median_price"`
}

func GetPriceInCurrency(appID, currency, marketHashName string) (*MarketPrice, error) {
	url := fmt.Sprintf(
		"https://steamcommunity.com/market/priceoverview/?appid=%s&currency=%s&market_hash_name=%s",
		appID, currency, marketHashName,
	)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var marketPrice MarketPrice
	err = json.Unmarshal(body, &marketPrice)
	if err != nil {
		return nil, err
	}

	if !marketPrice.Success {
		return nil, fmt.Errorf("failed to get price for currency %s", currency)
	}

	return &marketPrice, nil
}

func ExtractPrice(priceStr string) float64 {
	log.Printf("ExtractPrice input: %s", priceStr)
	re := regexp.MustCompile(`[\d.,]+`)
	parts := re.FindAllString(priceStr, -1)
	if len(parts) == 0 {
		log.Printf("No number parts found in: %s", priceStr)
		return 0
	}
	priceStr = parts[0]
	log.Printf("Extracted number part: %s", priceStr)

	dotCount := strings.Count(priceStr, ".")
	commaCount := strings.Count(priceStr, ",")

	if dotCount > 1 {
		lastIndex := strings.LastIndex(priceStr, ".")
		wholePart := priceStr[:lastIndex]
		decimalPart := priceStr[lastIndex+1:]
		wholePart = strings.ReplaceAll(wholePart, ".", "")
		priceStr = wholePart + "." + decimalPart
		log.Printf("After processing multiple dots: %s", priceStr)
	} else if commaCount > 1 {
		lastIndex := strings.LastIndex(priceStr, ",")
		wholePart := priceStr[:lastIndex]
		decimalPart := priceStr[lastIndex+1:]
		wholePart = strings.ReplaceAll(wholePart, ",", "")
		priceStr = wholePart + "." + decimalPart
		log.Printf("After processing multiple commas: %s", priceStr)
	} else if dotCount == 1 && commaCount == 1 {
		dotIndex := strings.Index(priceStr, ".")
		commaIndex := strings.Index(priceStr, ",")
		if commaIndex > dotIndex {
			priceStr = strings.ReplaceAll(priceStr, ".", "")
			priceStr = strings.Replace(priceStr, ",", ".", 1)
		} else {
			priceStr = strings.ReplaceAll(priceStr, ",", "")
		}
		log.Printf("After processing mixed separators: %s", priceStr)
	} else {
		priceStr = strings.Replace(priceStr, ",", ".", 1)
		log.Printf("After replacing comma with dot: %s", priceStr)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		log.Printf("Error parsing price: %v from string %s", err, priceStr)
		return 0
	}
	log.Printf("Parsed price: %f", price)
	return price
}


var currenciesInCents = map[string]bool{

}

type CurrencyResult struct {
	Name  string
	Rate  float64
	Error error
}

func GetSteamCurrencyRates(appID, marketHashName string) (map[string]float64, error) {
	currencies := map[string]string{
		"USD": "1",  // United States Dollar
		"EUR": "3",  // European Union Euro
		"RUB": "5",  // Russian Rouble
		"JPY": "8",  // Japanese Yen
		"CNY": "23", // Chinese Renminbi (yuan)
	}

	baseCurrencyCode := currencies["USD"]
	basePrice, err := GetPriceInCurrency(appID, baseCurrencyCode, marketHashName)
	if err != nil {
		return nil, err
	}
	if basePrice.LowestPrice == "" {
		return nil, fmt.Errorf("no lowest price for base currency (USD)")
	}
	basePriceFloat := ExtractPrice(basePrice.LowestPrice)

	if basePriceFloat == 0 {
		return nil, fmt.Errorf("base price is zero, cannot calculate rates")
	}

	semaphore := make(chan struct{}, 5)
	results := make(chan CurrencyResult, len(currencies))
	var wg sync.WaitGroup

	for name, code := range currencies {
		wg.Add(1)
		go func(name, code string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			time.Sleep(200 * time.Millisecond)

			price, err := GetPriceInCurrency(appID, code, marketHashName)
			if err != nil {
				results <- CurrencyResult{Name: name, Rate: 0, Error: err}
				return
			}
			if price.LowestPrice == "" {
				results <- CurrencyResult{Name: name, Rate: 0, Error: fmt.Errorf("no price")}
				return
			}
			priceFloat := ExtractPrice(price.LowestPrice)

			if currenciesInCents[code] {
				priceFloat /= 100
			}

			rate := math.Trunc((priceFloat/basePriceFloat)*100) / 100
			results <- CurrencyResult{Name: name, Rate: rate, Error: nil}
		}(name, code)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	rates := make(map[string]float64)

	for res := range results {
		if res.Error != nil {
			log.Printf("Failed to get price for %s: %v", res.Name, res.Error)
			continue
		}
		rates[res.Name] = res.Rate
	}

	return rates, nil
}