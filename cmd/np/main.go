package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

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

	paymentID := flag.String("p", "", "status of payment ID")
	newPayment := flag.Bool("n", false, "new payment")
	payAmount := flag.Float64("a", 2.0, "pay amount for new payment")
	flag.Parse()

	if *paymentID != "" {
		ps, err := payments.Status(*paymentID)
		if err != nil {
			log.Fatal(err)
		}
		d, err := json.Marshal(ps)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(d))
		return
	}

	fmt.Println("Sandbox:", core.BaseURL() == core.SandBoxBaseURL)
	st, err := core.Status()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("API Status:", st)

	cs, err := currencies.All()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d available crypto currencies\n", len(cs))

	if *newPayment {
		pa := &payments.PaymentArgs{
			PaymentAmount: payments.PaymentAmount{
				PriceAmount:      *payAmount,
				PriceCurrency:    "eur",
				PayCurrency:      "xmr",
				OrderID:          "tool 1",
				OrderDescription: "Some useful tool",
			},
		}
		fmt.Printf("Creating a %.2f payment ...\n", pa.PriceAmount)
		pay, err := payments.New(pa)
		if err != nil {
			log.Fatal(err)
		}
		d, err := json.Marshal(pay)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(d))
	}
}
