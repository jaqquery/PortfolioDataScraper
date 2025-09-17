package main

import (
	"fmt"
)

// "database/sql"
// "fmt"
// "log"
// "time"
// _ "github.com/go-sql-driver/mysql"

// type Cashflow struct {
// 	ID        int
// 	Portfolio string
// 	Value     float64
// }

// type DbCredentials struct {
// 	dbUser     string // your MariaDB username
// 	dbPassword string // your MariaDB password
// 	dbHost     string // or "127.0.0.1"
// 	dbPort     string // default MariaDB port
// 	dbName     string // your database name
// }

// func dbConnection() (*sql.Rows, error) {
// 	var credit DbCredentials
// 	credit.dbUser = "root"
// 	credit.dbPassword = "Cashflow5426#"
// 	credit.dbHost = "localhost"
// 	credit.dbPort = "3306"
// 	credit.dbName = "cashflow"

// 	// Create connection string
// 	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
// 		credit.dbUser, credit.dbPassword, credit.dbHost, credit.dbPort, credit.dbName)

// 	// Open database connection
// 	db, err := sql.Open("mysql", connectionString)
// 	if err != nil {
// 		log.Fatal("Error opening database:", err)
// 	}
// 	defer db.Close()

// 	// Test the connection
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("Error connecting to database:", err)
// 	}
// 	fmt.Println("Successfully connected to MariaDB!")

// 	// Query to read from your table
// 	// Replace "your_table" with your actual table name
// 	query := "SELECT id, portfolio, value FROM cashflow"

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Fatal("Error executing query:", err)
// 	}
// 	defer rows.Close()

// 	return rows, nil

// }

//TODO: CREATE crytocurrency, retirement table
//THIS SERVICES will keep running to update latest price available for crypto/stocks portfolio

func main() {

	price, err := CRYPTO_PRICE_FROM_WEB("bitcoin")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Bitcoin price: $%.4f\n", price)
	}

	price, err = CRYPTO_PRICE_FROM_WEB("ethereum")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Ethereum price: $%.4f\n", price)
	}

	//TODO: Alter table, add column "countrycode"
	//maybe just use one table for the different country stocks
	price, name, err := getStockPrice("5227", "MY")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Stock name: %s, Regular Market Price: %.2f\n", name, price)

}
