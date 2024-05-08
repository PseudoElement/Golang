package oneinch

import (
	"fmt"
	"net/http"
)

func (m *OneinchModule) SetRoutes() {
	m.router.HandleFunc("/oneinch/quote", m._quoteController).Methods(http.MethodGet)
	m.router.HandleFunc("/oneinch/swap", m._swapController).Methods(http.MethodGet)
	m.router.HandleFunc("/oneinch/allowance", m._getAllowanceController).Methods(http.MethodGet)
	m.router.HandleFunc("/oneinch/approve", m._getApproveConfigController).Methods(http.MethodGet)
	m.router.HandleFunc("/hello", m._helloController).Methods(http.MethodGet)

	fmt.Println("OneinchModule started!")
}
