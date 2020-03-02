package main

import (
	"flag"
	"fmt"
	"github.com/centum/blockonomics"
	"os"
	"time"
)

func main() {
	addrCommand := flag.NewFlagSet("addr", flag.ExitOnError)
	token := addrCommand.String("token", "", "access token to API Blockonomics")
	account := addrCommand.String("account", "", "get address for account")
	reset := addrCommand.Bool("reset", false, "reset prev address")

	invoiceCommand := flag.NewFlagSet("invoice", flag.ExitOnError)
	addr := invoiceCommand.String("addr", "", "Invoice for address")
	amount := invoiceCommand.Float64("amount", 0, "Invoice amount")
	currency := invoiceCommand.String("currency", "USD", "Invoice currency")
	description := invoiceCommand.String("description", "", "Invoice description")
	invoiceLive := invoiceCommand.Duration("live", 1*time.Hour, "Invoice live time")
	flag.Parse()

	switch {
	case len(os.Args) > 1 && os.Args[1] == "addr":
		addrCommand.Parse(os.Args[2:])
		fmt.Printf("==> REQUEST:\n  token=%s\n  account=%s\n  reset=%v\n==> RESPONSE:\n", *token, *account, *reset)

		api := blockonomics.NewClient(*token, blockonomics.WithTimeout(30*time.Second))
		a, err := api.NewAddress("", true)
		if err != nil {
			panic(err)
		}
		fmt.Printf("addr: %s\naccount: %s\nreset: %v", a.Address, a.Account, a.Reset == 1)

	case len(os.Args) > 1 && os.Args[1] == "invoice":
		invoiceCommand.Parse(os.Args[2:])
		fmt.Printf(
			"==> REQUEST:\n  addr=%s\n  amount=%.2f\n  currency=%s\n  description=%s\n  live=%v\n==> RESPONSE:",
			*addr, *amount, *currency, *description, *invoiceLive)

		api := blockonomics.NewClient("", blockonomics.WithTimeout(30*time.Second))
		checkoutURL, err := api.Invoice(*addr, *amount, *currency, *description, time.Now().Add(*invoiceLive))
		if err != nil {
			panic(err)
		}
		fmt.Printf("  checkoutURL: %s\n", checkoutURL)

	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		fmt.Printf("usage: %s <command> [<args>]", os.Args[0])
		os.Exit(-1)
	}
}
