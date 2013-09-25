package server

type IWebApplication interface {
	Start()
	ReStart()
	Stop()
}
