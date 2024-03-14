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
	// walletAddress := req.URL.Query().Get("walletAddress");

	quoteUrl := fmt.Sprintf("%v/%v/quote", oneinch_api_url, "56");

	res, err := http.Get(quoteUrl);

	if err != nil {
		w.WriteHeader(http.StatusBadRequest);
		w.Write([]byte("Invalid params!"));
		return;
	}

	quoteResBody, _ := io.ReadAll(res.Body);
	quoteResBodyString := string(quoteResBody);

	fmt.Println("BODY_RAW ", quoteResBodyString);

	w.Header().Set("Content-Type", "application/json");
	w.Header().Set("Authorization", oneinch_authorization_header_value);

	w.WriteHeader(http.StatusOK);
	json.NewEncoder(w).Encode(quoteResBodyString)
}