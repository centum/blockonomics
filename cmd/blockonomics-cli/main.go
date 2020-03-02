package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/centum/blockonomics"
	"os"
	"time"
)

func main() {
	switch {
	case len(os.Args) > 1 && os.Args[1] == "balance":
		cmd := flag.NewFlagSet("balance", flag.ExitOnError)
		addr := cmd.String("addr", "", "Whitespace separated list of bitcoin addresses/xpubs")
		cmd.Parse(os.Args[2:])

		api := blockonomics.NewClient("", blockonomics.WithTimeout(30*time.Second))
		dump(api.Balance(*addr))

	case len(os.Args) > 1 && os.Args[1] == "searchhistory":
		cmd := flag.NewFlagSet("searchhistory", flag.ExitOnError)
		addr := cmd.String("addr", "", "Whitespace separated list of bitcoin addresses/xpubs")
		cmd.Parse(os.Args[2:])

		api := blockonomics.NewClient("", blockonomics.WithTimeout(30*time.Second))
		dump(api.SearchHistory(*addr))

	case len(os.Args) > 1 && os.Args[1] == "tx_detail":
		cmd := flag.NewFlagSet("tx_detail", flag.ExitOnError)
		txid := cmd.String("txid", "", "transaction id")
		cmd.Parse(os.Args[2:])

		api := blockonomics.NewClient("", blockonomics.WithTimeout(30*time.Second))
		dump(api.TxDetail(*txid))

	case len(os.Args) > 1 && os.Args[1] == "addr":
		cmd := flag.NewFlagSet("addr", flag.ExitOnError)
		token := cmd.String("token", "", "access token to API Blockonomics")
		account := cmd.String("account", "", "get address for account")
		reset := cmd.Bool("reset", false, "reset prev address")
		cmd.Parse(os.Args[2:])

		api := blockonomics.NewClient(*token, blockonomics.WithTimeout(30*time.Second))
		dump(api.NewAddress(*account, *reset))

	case len(os.Args) > 1 && os.Args[1] == "invoice":
		cmd := flag.NewFlagSet("invoice", flag.ExitOnError)
		addr := cmd.String("addr", "", "Invoice for address")
		amount := cmd.Float64("amount", 0, "Invoice amount")
		currency := cmd.String("currency", "USD", "Invoice currency")
		description := cmd.String("description", "", "Invoice description")
		invoiceLive := cmd.Duration("live", 1*time.Hour, "Invoice live time")
		cmd.Parse(os.Args[2:])

		api := blockonomics.NewClient("", blockonomics.WithTimeout(30*time.Second))
		dump(api.Invoice(*addr, *amount, *currency, *description, time.Now().Add(*invoiceLive)))

	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		fmt.Printf("usage: %s <command> [<args>]", os.Args[0])
		os.Exit(-1)
	}
}

func dump(v interface{}, err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
	d, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(d))
}
