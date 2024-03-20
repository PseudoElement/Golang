package OneinchModels

type QuoteRes struct {
	ToAmount string `json:"dstAmount"`
}

type SwapRes struct {
	ToAmount string      `json:"dstAmount"`
	Tx       Transaction `json:"tx"`
}

type Transaction struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Data  string `json:"data"`
	Value string `json:"value"`
	Gas   int    `json:"gas"`
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
