package main

import (
	"fmt"
	"sync"

	"github.com/amsem/crypto/api"
)

func main() {
	var wg sync.WaitGroup
	currencies := []string{"BTC", "BCH", "ETH", "MATIC", "USDT"}
	for _, c := range currencies {
		wg.Add(1)
		go func(cur string) {
			getCurency(cur)
			wg.Done()
		}(c)
	}
	wg.Wait()
}

func getCurency(c string) {
	rate, err := api.GetRate(c)
	if err == nil {
		fmt.Printf("The currecy is %v and the price is %.6f\n", rate.Currency, rate.Price)
	}

}
