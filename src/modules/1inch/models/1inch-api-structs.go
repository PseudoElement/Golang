package OneinchModels

type QuoteRes struct {
	ToAmount string `json:"dstAmount"`
}
type GetSpenderAddressRes struct {
	SpenderAddress string `json:"address"`
}

type GetTokenAllowanceRes struct {
	Allowance string `json:"allowance"`
}

type GetApproveConfigRes struct {
	Data     string `json:"data"`
	To       string `json:"to"`
	Value    string `json:"value"`
	GasPrice string `json:"gasPrice"`
}
