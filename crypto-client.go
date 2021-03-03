package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type CoinValue struct {
	Usd            float32
	Usd_Market_Cap float32
}

func GetCryptoUpdates() {
	tick := time.NewTicker(time.Second * 3)
	go GetCryptoPrices(tick)

	select {}
}

func GetCryptoPrices(ticker *time.Ticker) {
	for range ticker.C {
		parameters := make(map[string]string)
		parameters["ids"] = `bitcoin%2Cethereum`
		parameters["vs_currencies"] = "usd"
		parameters["include_market_cap"] = "true"
		// max 100 calls per minute in the api
		// curl -X GET "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd&include_market_cap=true
		const coinPriceUrl = "https://api.coingecko.com/api/v3/simple/price?"
		url := BuildApiUrl(coinPriceUrl, parameters)

		response, err := http.Get(url)
		if err != nil {
			color.Red("Error getting info from ", url)
		} else {
			// https://stackoverflow.com/questions/17452722/how-to-get-the-key-value-from-a-json-string-in-go
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err == nil {
				jsonData := make(map[string]json.RawMessage)
				err := json.Unmarshal(body, &jsonData)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(body))
				coins := make(map[string]CoinValue)
				i := 0
				for keyValue, value := range jsonData {
					var coinValue = new(CoinValue)
					json.Unmarshal(value, coinValue)
					coins[keyValue] = *coinValue
					// fmt.Println(string(value))
					i++
				}
				PrintAllCoins(coins)
			} else {
				color.Red("Error reading the body")
			}
		}
	}
}

func BuildApiUrl(baseUrl string, parameters map[string]string) string {
	parameterString := ""
	delimeter := "&"
	for key, value := range parameters {
		parameterString += (key + "=" + value + delimeter)
	}
	parameterString = parameterString[:len(parameterString)-1]
	fmt.Println(baseUrl + parameterString)
	return (baseUrl + parameterString)
}

func PrintAllCoins(coins map[string]CoinValue) {
	for key := range coins {
		color.HiGreen(key + "\n")
		color.HiMagenta(GetCoinValue(coins[key]))
	}
}

func GetCoinValue(coin CoinValue) string {
	return fmt.Sprintf("USD: $%.2f Market Cap USD: $%.2f", coin.Usd, coin.Usd_Market_Cap)
}
