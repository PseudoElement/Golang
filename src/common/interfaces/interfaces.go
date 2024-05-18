package interfaces_module

import (
	"github.com/gorilla/websocket"
	errors_module "github.com/pseudoelement/go-server/src/errors"
)

type ModuleWithRoutes interface {
	SetRoutes()
}

type Socket interface {
	Connect() errors_module.ErrorWithStatus
	Conn() *websocket.Conn
	Disconnect() errors_module.ErrorWithStatus
	Broadcast()
}
