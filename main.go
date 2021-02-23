package main

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
)

func main() {
	// get prices from exchanges of different cypto coin prices
	// us go routine do get the info in the background and pipe the infor to the display
	// write this info to kafka and then have the frontend get the data from kafka
	// somehow use kubernetes with this
	// run a service that logs this information up to a certain point
	// this service run along the main service to get the crypto prices
	PrintWelcomeScreen()
	GetCryptoPrices()
}

func PrintWelcomeScreen() {
	fmt.Println("Welcome to the Crypto Currency Tracker")
	fmt.Println("--------------------------------------")
	//	color.Cyan("This is Cyan")
	//	color.Red("Red is dead")

}

func GetCryptoPrices() {
	// get the current btc price here
	// https://api.coingecko.com/api/v3/ping
	var url = "https://api.coingecko.com/api/v3/ping"
	response, err := http.Get(url)
	if err != nil {
		color.Red("Error getting info from ", url)
	} else {
		fmt.Println("respone = ", response)
	}
}
