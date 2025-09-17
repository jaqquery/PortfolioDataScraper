package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

func CRYPTO_PRICE_FROM_WEB(coinSlug string) (float64, error) {
	// Fetch data from CoinMarketCap website
	url := "https://coinmarketcap.com/currencies/" + coinSlug + "/"
	response, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return 0, fmt.Errorf("received non-200 status code: %d", response.StatusCode)
	}

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %v", err)
	}

	contentStr := string(content)

	// Try different regex patterns to find the price
	price, err := extractPrice(contentStr)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func extractPrice(content string) (float64, error) {
	// Pattern 1: Look for price in JSON data
	priceRegex1 := regexp.MustCompile(`"price":\s*(\d+(\.\d+)?)`)
	match1 := priceRegex1.FindStringSubmatch(content)
	if len(match1) > 1 {
		price, err := strconv.ParseFloat(match1[1], 64)
		if err == nil {
			return price, nil
		}
	}

	// Pattern 2: Look for price in a different format
	priceRegex2 := regexp.MustCompile(`data-price="(\d+(\.\d+)?)"`)
	match2 := priceRegex2.FindStringSubmatch(content)
	if len(match2) > 1 {
		price, err := strconv.ParseFloat(match2[1], 64)
		if err == nil {
			return price, nil
		}
	}

	// Pattern 3: Another possible format
	priceRegex3 := regexp.MustCompile(`"priceValue"[^>]*>([^<]+)`)
	match3 := priceRegex3.FindStringSubmatch(content)
	if len(match3) > 1 {
		// Remove any currency symbols and commas
		cleanPattern := regexp.MustCompile(`[$,]`)
		rawPrice := cleanPattern.ReplaceAllString(match3[1], "")
		price, err := strconv.ParseFloat(rawPrice, 64)
		if err == nil {
			return price, nil
		}
	}

	return 0, fmt.Errorf("could not find or parse price information")
}

// func main() {
// 	// Example usage
// 	price, err := CRYPTO_PRICE_FROM_WEB("bitcoin")
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 	} else {
// 		fmt.Printf("Bitcoin price: $%.4f\n", price)
// 	}

// 	price, err = CRYPTO_PRICE_FROM_WEB("ethereum")
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 	} else {
// 		fmt.Printf("Ethereum price: $%.4f\n", price)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"math/rand"
// 	"net/http"
// 	"time"
// )

// func getCryptoPrice(symbol string) (float64, error) {

// 	url := fmt.Sprintf("https://coinmarketcap.com/currencies/%s", symbol)

// 	client := &http.Client{
// 		Timeout: 15 * time.Second,
// 	}

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Add some delay to avoid rate limiting
// 	time.Sleep(time.Duration(500+rand.Intn(500)) * time.Millisecond)

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return 1, fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return 0, err
// 	}

// 	fmt.Println("%s", body)

// 	return 0, nil

// }
