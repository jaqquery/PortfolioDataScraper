package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

//For MY stocks: 5227.KL
//For US stocks: AAPL
//For HK stocks: 1398.HK
//For SG stocks: ES3.SI

func getStockPrice(symbol string, countryCode string) (float64, string, error) {

	switch countryCode {
	case "MY":
		symbol = symbol + ".KL"
	case "HK":
		symbol = symbol + ".HK"
	case "SG":
		symbol = symbol + ".SI"
	case "US":
	default:
		symbol = symbol
	}

	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s", symbol)

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, "", err
	}

	// Set realistic browser headers
	req.Header = http.Header{
		"User-Agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"},
		"Accept":          {"application/json, text/plain, */*"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"Referer":         {"https://finance.yahoo.com/"},
		"Origin":          {"https://finance.yahoo.com"},
		"Connection":      {"keep-alive"},
		"Sec-Fetch-Dest":  {"empty"},
		"Sec-Fetch-Mode":  {"cors"},
		"Sec-Fetch-Site":  {"same-site"},
	}

	// Add some delay to avoid rate limiting
	time.Sleep(time.Duration(500+rand.Intn(500)) * time.Millisecond)

	resp, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 1, "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}

	var result struct {
		Chart struct {
			Result []struct {
				Meta struct {
					RegularMarketPrice float64 `json:"regularMarketPrice"`
					ShortName          string  `json:"shortname"`
				} `json:"meta"`
			} `json:"result"`
		} `json:"chart"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, "", err
	}

	if len(result.Chart.Result) == 0 {
		return 0, "", fmt.Errorf("no data found for symbol %s", symbol)
	}

	return result.Chart.Result[0].Meta.RegularMarketPrice, result.Chart.Result[0].Meta.ShortName, nil
}
