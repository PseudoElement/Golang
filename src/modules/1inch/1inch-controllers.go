package Oneinch

import (
	"encoding/json"
	"fmt"
	ApiService "go-server/src/api"
	"net/http"
)
func _helloController(w http.ResponseWriter, req *http.Request){
	count++;

	w.WriteHeader(http.StatusAccepted);
	fmt.Fprintf(w, "Hello controller calls %v times!\n", count);
}

func _quoteController(w http.ResponseWriter, req *http.Request) {
	ApiService.SetResponseHeaders(w, req);
	
	quoteParams := ApiService.MapQueryParams(req, "src", "dst", "amount", "chainId");
	quoteHeaders := map[string]string{"Authorization": oneinch_authorization_header_value};
	quoteUrl := fmt.Sprintf("%v/%v/quote", oneinch_api_url, quoteParams["chainId"]);

	res, err := ApiService.Get(quoteUrl, quoteParams, quoteHeaders)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest);
		w.Write([]byte("Invalid params!"));
		return;
	}

	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(res);
}

func _swapController(w http.ResponseWriter, req *http.Request) {
	ApiService.SetResponseHeaders(w, req);

	swapParams := ApiService.MapQueryParams(req, "src", "dst", "amount", "chainId", "from", "receiver", "slippage");
	swapHeaders := map[string]string{"Authorization": oneinch_authorization_header_value};
	swapUrl := fmt.Sprintf("%v/%v/swap", oneinch_api_url, swapParams["chainId"]);

	res, err := ApiService.Get(swapUrl, swapParams, swapHeaders)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest);
		w.Write([]byte("Invalid params!"));
		return;
	}

	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(res);
}