package interfaces

type ModuleWithRoutes interface {
	SetRoutes()
}

type Socket interface {
	Connect()
	Disconnect()
	Listen()
}
