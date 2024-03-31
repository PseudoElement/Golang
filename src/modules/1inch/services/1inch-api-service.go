package OneinchApiService

import (
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

	resBody, resErr := ApiService.Get[OneinchModels.QuoteRes](quoteUrl, quoteParams, quoteHeaders)
	if resErr != nil {
		return OneinchModels.QuoteRes{}, resErr;
	}

	return resBody, nil;
}

func MakeSwapRequest(w http.ResponseWriter, req *http.Request) (OneinchModels.SwapRes, error) {
	params := ApiService.MapQueryParams(req, "src", "dst", "amount", "chainId", "from", "receiver", "slippage");
	headers := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE};
	url := fmt.Sprintf("%v/%v/swap", OneinchConsts.ONEINCH_API_URL, params["chainId"]);

	resBody, resErr := ApiService.Get[OneinchModels.SwapRes](url, params, headers);
	if resErr != nil {
		return OneinchModels.SwapRes{}, resErr;
	}

	return resBody, nil;
}

func GetSpenderAddress(w http.ResponseWriter, chainId string) (OneinchModels.GetSpenderAddressRes, error) {
	headers := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE}
	url := fmt.Sprintf("%v/%v/approve/spender", OneinchConsts.ONEINCH_API_URL, chainId);

	resBody, resErr := ApiService.Get[OneinchModels.GetSpenderAddressRes](url, make(map[string]string), headers);
	if resErr != nil{
		return OneinchModels.GetSpenderAddressRes{}, resErr;
	}

	return resBody, nil;
}

func GetTokenAllowance(w http.ResponseWriter, req *http.Request) (OneinchModels.GetTokenAllowanceRes, error) {
	allQueryParams := ApiService.MapQueryParams(req, "src", "chainId", "walletAddress");
	headers := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE}
	url := fmt.Sprintf("%v/%v/approve/allowance", OneinchConsts.ONEINCH_API_URL, allQueryParams["chainId"]);
	params := map[string]string{"tokenAddress": allQueryParams["src"], "walletAddress": allQueryParams["walletAddress"]};

	resBody, resErr := ApiService.Get[OneinchModels.GetTokenAllowanceRes](url, params, headers);
	if resErr != nil{
		return OneinchModels.GetTokenAllowanceRes{}, resErr;
	}

	return resBody, nil;
}

func GetApproveConfig(w http.ResponseWriter, chainId string, fromTokenAddress string, fromTokenAmount string)  (OneinchModels.GetApproveConfigRes, error) {
	headers := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE};
	params := map[string]string{"tokenAddress": fromTokenAddress, "amount": fromTokenAmount};
	url := fmt.Sprintf("%v/%v/approve/transaction", OneinchConsts.ONEINCH_API_URL, chainId);

	resBody, resErr := ApiService.Get[OneinchModels.GetApproveConfigRes](url, params, headers);
	if resErr != nil{
		return OneinchModels.GetApproveConfigRes{}, resErr;
	}

	return resBody, nil;
}