# PortfolioDataScraper
Backend services retrieving data from the internet

#For cryptocurrencies
USAGE:
price, err = CRYPTO_PRICE_FROM_WEB("ethereum")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Ethereum price: $%.4f\n", price)
	}

#For MY, HK and SI Stockmarket
USAGE
price, name, err := getStockPrice("5227", "MY")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

 #Pending:
 US Market


 
