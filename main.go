package main

import (
	"fmt"
)

func main() {
	// get prices from exchanges of different cypto coin prices
	// us go routine do get the info in the background and pipe the infor to the display
	// write this info to kafka and then have the frontend get the data from kafka
	// somehow use kubernetes with this
	// run a service that logs this information up to a certain point
	// this service run along the main service to get the crypto prices
	PrintWelcomeScreen()
	GetCryptoUpdates()
}

func PrintWelcomeScreen() {
	fmt.Println("Welcome to the Crypto Currency Tracker")
	fmt.Println("--------------------------------------")
}
