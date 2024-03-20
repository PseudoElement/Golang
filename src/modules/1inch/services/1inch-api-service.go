package OneinchApiService

import (
	"encoding/json"
	"fmt"
	ApiService "go-server/src/api"
	OneinchConsts "go-server/src/modules/1inch/constants"
	OneinchModels "go-server/src/modules/1inch/models"
	"net/http"
)


func MakeQuoteRequest(w http.ResponseWriter, req *http.Request) (OneinchModels.QuoteRes, error) {
	quoteParams := ApiService.MapQueryParams(req, "src", "dst", "amount", "chainId");
	quoteHeaders := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE};
	quoteUrl := fmt.Sprintf("%v/%v/quote", OneinchConsts.ONEINCH_API_URL, quoteParams["chainId"]);

	resBody, resErr := ApiService.Get(quoteUrl, quoteParams, quoteHeaders)

	if resErr != nil {
		return OneinchModels.QuoteRes{}, resErr;
	}

	var resObject OneinchModels.QuoteRes;
	if err := json.Unmarshal(resBody, &resObject); err != nil {
		return OneinchModels.QuoteRes{}, err;
	}

	return resObject, nil;
}

func MakeSwapRequest(w http.ResponseWriter, req *http.Request) (OneinchModels.SwapRes, error) {
	params := ApiService.MapQueryParams(req, "src", "dst", "amount", "chainId", "from", "receiver", "slippage");
	headers := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE};
	url := fmt.Sprintf("%v/%v/swap", OneinchConsts.ONEINCH_API_URL, params["chainId"]);

	resBody, resErr := ApiService.Get(url, params, headers);

	if resErr != nil {
		return OneinchModels.SwapRes{}, resErr;
	}

	var swapStruct OneinchModels.SwapRes;
	if err := json.Unmarshal(resBody, &swapStruct); err != nil {
		return OneinchModels.SwapRes{}, err;
	}

	return swapStruct, nil;
}

func GetSpenderAddress(w http.ResponseWriter, chainId string) (OneinchModels.GetSpenderAddressRes, error) {
	headers := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE}
	url := fmt.Sprintf("%v/%v/approve/spender", OneinchConsts.ONEINCH_API_URL, chainId);

	resBody, resErr := ApiService.Get(url, make(map[string]string), headers);
	if resErr != nil{
		return OneinchModels.GetSpenderAddressRes{}, resErr;
	}

	var approveStruct OneinchModels.GetSpenderAddressRes
	if err := json.Unmarshal(resBody, &approveStruct); err != nil {
		return OneinchModels.GetSpenderAddressRes{}, err;
	}

	return approveStruct, nil;
}

func GetTokenAllowance(w http.ResponseWriter, req *http.Request) (OneinchModels.GetTokenAllowanceRes, error) {
	allQueryParams := ApiService.MapQueryParams(req, "src", "chainId", "walletAddress");
	headers := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE}
	url := fmt.Sprintf("%v/%v/approve/allowance", OneinchConsts.ONEINCH_API_URL, allQueryParams["chainId"]);
	params := map[string]string{"tokenAddress": allQueryParams["src"], "walletAddress": allQueryParams["walletAddress"]};

	resBody, resErr := ApiService.Get(url, params, headers);
	if resErr != nil{
		return OneinchModels.GetTokenAllowanceRes{}, resErr;
	}

	var allowanceStruct OneinchModels.GetTokenAllowanceRes;
	if err := json.Unmarshal(resBody, &allowanceStruct); err != nil {
		return OneinchModels.GetTokenAllowanceRes{}, err;
	}

	return allowanceStruct, nil;
}

func GetApproveConfig(w http.ResponseWriter, chainId string, fromTokenAddress string, fromTokenAmount string)  (OneinchModels.GetApproveConfigRes, error) {
	headers := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE};
	params := map[string]string{"tokenAddress": fromTokenAddress, "amount": fromTokenAmount};
	url := fmt.Sprintf("%v/%v/approve/transaction", OneinchConsts.ONEINCH_API_URL, chainId);

	resBody, resErr := ApiService.Get(url, params, headers);
	if resErr != nil{
		return OneinchModels.GetApproveConfigRes{}, resErr;
	}

	var approveConfig OneinchModels.GetApproveConfigRes;
	if err := json.Unmarshal(resBody, &approveConfig); err != nil{
		return OneinchModels.GetApproveConfigRes{}, resErr;
	}

	return approveConfig, nil;
}