package OneinchApiService

import (
	"encoding/json"
	"fmt"
	ApiService "go-server/src/api"
	OneinchConsts "go-server/src/modules/1inch/constants"
	OneinchModels "go-server/src/modules/1inch/models"
	"net/http"
)


func GetApproveAddress(w http.ResponseWriter, chainId string) (approveAddress string, e error) {
	headers := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE}
	url := fmt.Sprintf("%v/%v/approve/spender", OneinchConsts.ONEINCH_API_URL, chainId);

	resBody, resErr := ApiService.Get(url, make(map[string]string), headers);
	if resErr != nil{
		return "", resErr;
	}

	var approveStruct OneinchModels.GetApproveAddressRes
	if err := json.Unmarshal(resBody, &approveStruct); err != nil {
		return "", err;
	}

	return approveStruct.ApproveAddress, nil;
}

func GetTokenAllowance(w http.ResponseWriter, chainId string) (allowance string, e error) {
	headers := map[string]string{"Authorization": OneinchConsts.ONEINCH_AUTHORIZATION_HEADER_VALUE}
	url := fmt.Sprintf("%v/%v/approve/allowance", OneinchConsts.ONEINCH_API_URL, chainId);

	resBody, resErr := ApiService.Get(url, make(map[string]string), headers);
	if resErr != nil{
		return "", resErr;
	}

	var allowanceStruct OneinchModels.GetTokenAllowanceRes;
	if err := json.Unmarshal(resBody, &allowanceStruct); err != nil {
		return "", err;
	}

	return allowanceStruct.Allowance, nil;
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