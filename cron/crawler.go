package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"strconv"
	"github.com/robfig/cron"
)

type EthMarketPriceOfMax struct {
	At uint `json:"at"`
	Buy string `json:"buy"`
	Sell string `json:"sell"`
	Open string `json:"open"`
	Low string `json:"low"`
	High string `json:"high"`
	Last string `json:"last"`
	Vol string `json:"vol"`
}

func getPrice() *EthMarketPriceOfMax {
	// max exchange api
	resp, err := http.Get("https://max-api.maicoin.com/api/v2/tickers/ethtwd")
	if err != nil {
		log.Println("http error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	price := new(EthMarketPriceOfMax)
	err = json.Unmarshal(body, &price)
	if err != nil {
		// handle error
	}
	return price
}

func main() {
	c := cron.New()
	lastStrikePrice := 0.0
	spec := "* * * * * *"
	c.AddFunc(spec, func() {
		ethPrice := getPrice()

		latestPrice, err := strconv.ParseFloat(ethPrice.Last, 64)
		if err != nil {
			// handle error
		}

		if lastStrikePrice != latestPrice && ethPrice.At != 0 {
			lastStrikePrice = latestPrice
			fmt.Printf("Timestamp: %d, Ether price: %s\n", ethPrice.At, ethPrice.Last)
		}
	})
	c.Start()
	select{}
	
}