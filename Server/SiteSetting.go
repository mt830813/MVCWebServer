package Server

type SiteSetting struct {
	Handlers         []*handlerSetting
	AppSetting       map[string]string
	ConnectionString map[string]string
}

func newSiteSetting() *SiteSetting {
	returnValue := new(SiteSetting)
	returnValue.AppSetting = make(map[string]string)
	returnValue.ConnectionString = make(map[string]string)
	returnValue.Handlers = make([]*handlerSetting, 0)

	return returnValue
}
