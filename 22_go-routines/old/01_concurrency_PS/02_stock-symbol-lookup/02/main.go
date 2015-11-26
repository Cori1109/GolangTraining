package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	for _, symbol := range stockSymbols {
		resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		quote := new(quoteResponse)
		xml.Unmarshal(body, &quote)

		fmt.Printf("%s: $%.2f\n", quote.Name, quote.LastPrice)
	}

	elapsed := time.Since(start)

	fmt.Printf("Execution Time: %s", elapsed)
}

type quoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}
