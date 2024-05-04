package interfaces_module

import errors_module "github.com/pseudoelement/go-server/src/errors"

type ModuleWithRoutes interface {
	SetRoutes()
}

type Socket interface {
	Connect()
	Disconnect() errors_module.ErrorWithStatus
	Broadcast()
}
