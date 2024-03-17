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

type ApproveResponse struct {
	ApproveAddress string `json:"address"`
}

type AllowanceResponse struct {
	Allowance string `json:"allowance"`
}
