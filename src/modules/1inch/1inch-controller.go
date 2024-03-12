package oneinch

import (
	"fmt"
	"net/http"
)


func QuoteController(w http.ResponseWriter, req *http.Request) {
	count++;
	walletAddress := req.URL.Query().Get("walletAddress")

	fmt.Printf("QuoteController triggered %v times. QueryParams are %v\n", count, walletAddress)
}