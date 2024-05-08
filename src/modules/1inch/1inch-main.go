package oneinch

import (
	"github.com/gorilla/mux"
	interfaces_module "github.com/pseudoelement/go-server/src/common/interfaces"
)

type OneinchModule struct {
	router *mux.Router
}

func NewModule(router *mux.Router) *OneinchModule {
	return &OneinchModule{router: router}
}

var _ interfaces_module.ModuleWithRoutes = (*OneinchModule)(nil)
