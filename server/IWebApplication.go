package Server

type IWebApplication interface {
	Start()
	ReStart()
	Stop()
}
