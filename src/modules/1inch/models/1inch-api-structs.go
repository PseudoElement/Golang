package OneinchModels

// type OneinchQuoteReqParams struct {
// 	src    string
// 	dst    string
// 	amount string
// }

// type OneinchQuoteResponse struct {
// 	toAmount  string
// 	gas       int
// 	fromToken interface{}
// 	toToken   interface{}
// 	protocols interface{}
// }

type GetApproveAddressRes struct {
	ApproveAddress string `json:"address"`
}

type GetTokenAllowanceRes struct {
	Allowance string `json:"allowance"`
}

type GetApproveConfigRes struct {
	Data string `json:"data"`
	To string `json:"to"`
	Value string `json:"value"`
	GasPrice string `json:"gasPrice"`
}
