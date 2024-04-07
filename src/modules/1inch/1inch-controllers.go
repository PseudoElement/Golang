package oneinch

import (
	"encoding/json"
	"fmt"
	"net/http"

	api_main "github.com/pseudoelement/go-server/src/api"
	oneinch_api "github.com/pseudoelement/go-server/src/modules/1inch/services"
)

func _helloController(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK);
	fmt.Fprintln(w, "Hello Controller triggered!")
}

func _getApproveConfigController(w http.ResponseWriter, req *http.Request) {
	api_main.SetResponseHeaders(w, req);

	params := api_main.MapQueryParams(req, "src", "amount", "chainId");

	approveObj, _ := oneinch_api.GetApproveConfig(w, params["chainId"], params["src"], params["amount"]);

	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(approveObj);
}

func _getAllowanceController(w http.ResponseWriter, req *http.Request) {
	api_main.SetResponseHeaders(w, req);

	allowanceObj, _ := oneinch_api.GetTokenAllowance(w, req);

	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(allowanceObj);
}

func _quoteController(w http.ResponseWriter, req *http.Request) {
	api_main.SetResponseHeaders(w, req);
	
	quoteData, _ := oneinch_api.MakeQuoteRequest(w, req);

	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(quoteData);
}

func _swapController(w http.ResponseWriter, req *http.Request) {
	api_main.SetResponseHeaders(w, req);

	swapData, err := oneinch_api.MakeSwapRequest(w, req);
	if err != nil {
		w.WriteHeader(http.StatusBadRequest);
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(swapData);
}