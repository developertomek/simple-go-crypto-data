package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/developertomek/crypto/api"
)

func main() {
	currencies := []string{"BTC", "ETH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func() {
			getCurrencyData(currency)
			wg.Done()
		}()
	}
	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
}
