package Oneinch

import (
	"encoding/json"
	"fmt"
	ApiService "go-server/src/api"
	OneinchApiService "go-server/src/modules/1inch/services"
	"net/http"
)

func _helloController(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK);
	fmt.Fprintln(w, "Hello Controller triggered!")
}

func _getApproveConfigController(w http.ResponseWriter, req *http.Request) {
	ApiService.SetResponseHeaders(w, req);

	params := ApiService.MapQueryParams(req, "src", "amount", "chainId");

	approveObj, _ := OneinchApiService.GetApproveConfig(w, params["chainId"], params["src"], params["amount"]);

	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(approveObj);
}

func _getAllowanceController(w http.ResponseWriter, req *http.Request) {
	ApiService.SetResponseHeaders(w, req);

	allowanceObj, _ := OneinchApiService.GetTokenAllowance(w, req);

	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(allowanceObj);
}

func _quoteController(w http.ResponseWriter, req *http.Request) {
	ApiService.SetResponseHeaders(w, req);
	
	quoteData, _ := OneinchApiService.MakeQuoteRequest(w, req);

	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(quoteData);
}

func _swapController(w http.ResponseWriter, req *http.Request) {
	ApiService.SetResponseHeaders(w, req);

	swapData, err := OneinchApiService.MakeSwapRequest(w, req);
	if err != nil {
		w.WriteHeader(http.StatusBadRequest);
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(swapData);
}