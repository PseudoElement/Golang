package oneinch

import (
	"encoding/json"
	"fmt"
	"io"
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


	client := &http.Client{};
	quoteUrl := fmt.Sprintf("%v/%v/quote", oneinch_api_url, chainId);
	quoteReq, err := http.NewRequest(http.MethodGet, quoteUrl, nil);
	
	params := quoteReq.URL.Query();
	params.Add("src", srcParam);
	params.Add("dst", dstParam);
	params.Add("amount", amountParam);
	
	quoteReq.URL.RawQuery = params.Encode();
	quoteReq.Header.Set("Authorization", oneinch_authorization_header_value);

	res, _ := client.Do(quoteReq);

	if err != nil {
		w.WriteHeader(http.StatusBadRequest);
		w.Write([]byte("Invalid params!"));
		return;
	}

	quoteResBody, _ := io.ReadAll(res.Body);
	quoteResBodyString := string(quoteResBody);

	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(quoteResBodyString);
}