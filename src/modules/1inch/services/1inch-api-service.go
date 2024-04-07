package oneinch_api

import (
	"fmt"
	"net/http"

	api_main "github.com/pseudoelement/go-server/src/api"
	oneinch_consts "github.com/pseudoelement/go-server/src/modules/1inch/constants"
	oneinch_models "github.com/pseudoelement/go-server/src/modules/1inch/models"
)


func MakeQuoteRequest(w http.ResponseWriter, req *http.Request) (oneinch_models.QuoteRes, error) {
	quoteParams := api_main.MapQueryParams(req, "src", "dst", "amount", "chainId");
	quoteHeaders := map[string]string{"Authorization": oneinch_consts.ONEINCH_AUTHORIZATION_HEADER_VALUE};
	quoteUrl := fmt.Sprintf("%v/%v/quote", oneinch_consts.ONEINCH_API_URL, quoteParams["chainId"]);

	resBody, resErr := api_main.Get[oneinch_models.QuoteRes](quoteUrl, quoteParams, quoteHeaders)
	if resErr != nil {
		return oneinch_models.QuoteRes{}, resErr;
	}

	return resBody, nil;
}

func MakeSwapRequest(w http.ResponseWriter, req *http.Request) (oneinch_models.SwapRes, error) {
	params := api_main.MapQueryParams(req, "src", "dst", "amount", "chainId", "from", "receiver", "slippage");
	headers := map[string]string{"Authorization": oneinch_consts.ONEINCH_AUTHORIZATION_HEADER_VALUE};
	url := fmt.Sprintf("%v/%v/swap", oneinch_consts.ONEINCH_API_URL, params["chainId"]);

	resBody, resErr := api_main.Get[oneinch_models.SwapRes](url, params, headers);
	if resErr != nil {
		return oneinch_models.SwapRes{}, resErr;
	}

	return resBody, nil;
}

func GetSpenderAddress(w http.ResponseWriter, chainId string) (oneinch_models.GetSpenderAddressRes, error) {
	headers := map[string]string{"Authorization": oneinch_consts.ONEINCH_AUTHORIZATION_HEADER_VALUE}
	url := fmt.Sprintf("%v/%v/approve/spender", oneinch_consts.ONEINCH_API_URL, chainId);

	resBody, resErr := api_main.Get[oneinch_models.GetSpenderAddressRes](url, make(map[string]string), headers);
	if resErr != nil{
		return oneinch_models.GetSpenderAddressRes{}, resErr;
	}

	return resBody, nil;
}

func GetTokenAllowance(w http.ResponseWriter, req *http.Request) (oneinch_models.GetTokenAllowanceRes, error) {
	allQueryParams := api_main.MapQueryParams(req, "src", "chainId", "walletAddress");
	headers := map[string]string{"Authorization": oneinch_consts.ONEINCH_AUTHORIZATION_HEADER_VALUE}
	url := fmt.Sprintf("%v/%v/approve/allowance", oneinch_consts.ONEINCH_API_URL, allQueryParams["chainId"]);
	params := map[string]string{"tokenAddress": allQueryParams["src"], "walletAddress": allQueryParams["walletAddress"]};

	resBody, resErr := api_main.Get[oneinch_models.GetTokenAllowanceRes](url, params, headers);
	if resErr != nil{
		return oneinch_models.GetTokenAllowanceRes{}, resErr;
	}

	return resBody, nil;
}

func GetApproveConfig(w http.ResponseWriter, chainId string, fromTokenAddress string, fromTokenAmount string)  (oneinch_models.GetApproveConfigRes, error) {
	headers := map[string]string{"Authorization": oneinch_consts.ONEINCH_AUTHORIZATION_HEADER_VALUE};
	params := map[string]string{"tokenAddress": fromTokenAddress, "amount": fromTokenAmount};
	url := fmt.Sprintf("%v/%v/approve/transaction", oneinch_consts.ONEINCH_API_URL, chainId);

	resBody, resErr := api_main.Get[oneinch_models.GetApproveConfigRes](url, params, headers);
	if resErr != nil{
		return oneinch_models.GetApproveConfigRes{}, resErr;
	}

	return resBody, nil;
}