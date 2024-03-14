package oneinch

type oneinchQuoteReqParams struct {
	src    string
	dst    string
	amount string
}

type oneinchQuoteResponse struct {
	toAmount  string
	gas       int
	fromToken interface{}
	toToken   interface{}
	protocols interface{}
}
