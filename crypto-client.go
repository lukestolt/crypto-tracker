package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
)

type CoinList struct {
	Bitcoin  Bitcoin
	Ethereum Ethereum
}

type Bitcoin struct {
	Bitcoin CoinValue
}

type Ethereum struct {
	Ethereum CoinValue
}

type CoinValue struct {
	Usd            int32
	Usd_Market_Cap float32
}

func GetCryptoPrices() {
	parameters := make(map[string]string)
	parameters["ids"] = `bitcoin%2Cethereum`
	parameters["vs_currencies"] = "usd"
	parameters["include_market_cap"] = "true"
	// curl -X GET "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd&include_market_cap=true"
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
			c := make(map[string]json.RawMessage)
			err := json.Unmarshal(body, &c)
			if err != nil {
				panic(err)
			}

			keys := make([]string, len(c))
			i := 0
			for keyValue, _ := range c {
				keys[i] = keyValue
				i++
			}

			fmt.Printf("%#v\n", keys)

			// var cValue CoinList
			// jsonError := json.Unmarshal(body, &cValue)
			// if jsonError != nil {
			// 	color.Red(jsonError.Error())
			// fmt.Println(cValue)
			//color.HiGreen(GetCoinValue(cValue))
			// fmt.Println(string(body)) // remember need to wrap in string otherwise its byte[]
		} else {
			color.Red("Error reading the body")
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

func GetCoinValue(coin Bitcoin) string {
	return fmt.Sprintf("USD: $%d Market Cap USD: $%.2f", coin.Bitcoin.Usd, coin.Bitcoin.Usd_Market_Cap)

}
