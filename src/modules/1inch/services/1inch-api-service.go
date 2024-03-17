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

	resBody, _ := ApiService.Get(url, make(map[string]string), headers);

	var approveStruct OneinchModels.OneinchApproveResponse
	if err := json.Unmarshal(resBody, &approveStruct); err != nil {
		return "", err;
	}

	return approveStruct.ApproveAddress, nil;
}