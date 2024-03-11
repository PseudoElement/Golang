package oneinchProxy

import (
	"fmt"
	"net/http"
)

func QuoteController(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	fmt.Println(query)
}