package oneinch

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
	w.Header().Set("Access-Control-Allow-Origin", "*");
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	srcParam := req.URL.Query().Get("src");
	dstParam := req.URL.Query().Get("dst");
	amountParam := req.URL.Query().Get("amount");
	chainId := req.URL.Query().Get("chainId")


	quoteUrl := fmt.Sprintf("%v/%v/quote", oneinch_api_url, chainId);
	quoteParams := map[string]string{"src": srcParam, "dst": dstParam, "amount": amountParam};
	quoteHeaders := map[string]string{"Authorization": oneinch_authorization_header_value};

	res, err := ApiService.Get(quoteUrl, quoteParams, quoteHeaders)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest);
		w.Write([]byte("Invalid params!"));
		return;
	}

	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(res);
}