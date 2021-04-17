package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/nerdthatnoonelikes/stonks/src/structs"
)

func main() {
	stonk := os.Args[1]

	// Makes a GET request to the API
	resp, err := http.Get("https://query1.finance.yahoo.com/v11/finance/quoteSummary/" + stonk + "?modules=summaryDetail,price")

	// Check for an error
	if err != nil {
		os.Exit(0)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var responseStruct structs.Response
	json.Unmarshal(bodyBytes, &responseStruct)

	symbol := responseStruct.Quotesummary.Result[0].Price.Symbol

	priceChange := responseStruct.Quotesummary.Result[0].Price.Regularmarketchange.Raw

	volume := responseStruct.Quotesummary.Result[0].Summarydetail.Volume.Fmt

	priceTraded := responseStruct.Quotesummary.Result[0].Summarydetail.Bid.Fmt

	percentChange := responseStruct.Quotesummary.Result[0].Price.Regularmarketchangepercent.Fmt

	regMarketChange := responseStruct.Quotesummary.Result[0].Price.Regularmarketchange.Fmt

	if priceChange < 0 {
		fmt.Printf(color.RedString("%s  %s   @  %s ﰬ %s   %s\n"), symbol, string(volume), string(priceTraded), string(regMarketChange), string(percentChange))
	} else if priceChange > 0 {
		fmt.Printf(color.GreenString("%s  %s   @  %s ﰵ %s   %s\n"), symbol, string(volume), string(priceTraded), string(regMarketChange), string(percentChange))
	}
}
