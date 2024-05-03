package interfaces_module

type ModuleWithRoutes interface {
	SetRoutes()
}

type Socket interface {
	Connect()
	Disconnect()
	Broadcast()
}
