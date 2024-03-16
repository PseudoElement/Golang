package Oneinch

type OneinchQuoteReqParams struct {
	src    string
	dst    string
	amount string
}

type OneinchQuoteResponse struct {
	toAmount  string
	gas       int
	fromToken interface{}
	toToken   interface{}
	protocols interface{}
}
