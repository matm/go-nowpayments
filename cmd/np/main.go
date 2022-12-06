package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/matm/go-nowpayments/pkg/config"
	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/matm/go-nowpayments/pkg/currencies"
	"github.com/matm/go-nowpayments/pkg/payments"
)

func main() {
	f, err := os.Open("conf.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = config.Load(f)
	if err != nil {
		log.Fatal(err)
	}
	core.UseBaseURL(core.SandBoxBaseURL)
	core.UseClient(core.NewHTTPClient())
	//core.WithDebug(true)

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

	ma, err := payments.MinimumAmount("xmr", "btc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Minimum amount:", ma.Amount)

	ps, err := payments.Status("abcd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Payment status:", ps)
}
