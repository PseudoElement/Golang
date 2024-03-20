package Oneinch

import (
	"encoding/json"
	"fmt"
	ApiService "go-server/src/api"
	OneinchConsts "go-server/src/modules/1inch/constants"
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

	params := ApiService.MapQueryParams(req, "src", "dst", "amount", "chainId", "from", "receiver", "slippage");
	headers := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE};
	swapUrl := fmt.Sprintf("%v/%v/swap", OneinchConsts.ONEINCH_API_URL, params["chainId"]);

	res, err := ApiService.Get(swapUrl, params, headers);
	//@TODO change on json object instead of string
	stringRes := string(res);

	if err != nil {
		w.WriteHeader(http.StatusBadRequest);
		w.Write([]byte("Invalid params!"));
		return;
	}

	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(stringRes);
}