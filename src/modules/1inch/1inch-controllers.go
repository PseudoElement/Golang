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

func _quoteController(w http.ResponseWriter, req *http.Request) {
	ApiService.SetResponseHeaders(w, req);
	
	quoteParams := ApiService.MapQueryParams(req, "src", "dst", "amount", "chainId");
	quoteHeaders := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE};
	quoteUrl := fmt.Sprintf("%v/%v/quote", OneinchConsts.ONEINCH_API_URL, quoteParams["chainId"]);

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

	params := ApiService.MapQueryParams(req, "src", "dst", "amount", "chainId", "from", "receiver", "slippage");
	headers := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE};
	swapUrl := fmt.Sprintf("%v/%v/swap", OneinchConsts.ONEINCH_API_URL, params["chainId"]);

	approveAddress, approveErr := OneinchApiService.GetApproveAddress(w, params["chainId"])

	if approveErr != nil{
		w.WriteHeader(http.StatusBadRequest);
		w.Write([]byte("Cannot get approve address!"));
		return;
	}

	res, err := ApiService.Get(swapUrl, params, headers);

	if err != nil {
		w.WriteHeader(http.StatusBadRequest);
		w.Write([]byte("Invalid params!"));
		return;
	}

	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(res);
}