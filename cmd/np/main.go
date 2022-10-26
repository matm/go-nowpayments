package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/matm/go-nowpayments/pkg/currencies"
	"github.com/matm/go-nowpayments/pkg/payments"
	"github.com/matm/go-nowpayments/pkg/types"
)

func main() {
	key := os.Getenv("NP_API_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "Missing API key")
		os.Exit(2)
	}
	core.UseAPIKey(key)
	core.UseBaseURL(types.ProductionBaseURL)

	st, err := core.Status()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("API Status:", st)

	cs, err := currencies.All()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Currencies:", cs)

	ep, err := payments.EstimatedPrice(100, "eur", "xmr")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Estimation: %f %s == %s %s\n", ep.AmountFrom, strings.ToUpper(ep.CurrencyFrom),
		ep.EstimatedAmount, strings.ToUpper(ep.CurrencyTo))
}
