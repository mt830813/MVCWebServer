package Server

type ViewModelBase struct {
	TopScript    string
	BottomScript string
	Css          string
}

func (this *ViewModelBase) SetTopScript(str string) {
	this.TopScript = str
}
func (this *ViewModelBase) SetBottomScript(str string) {
	this.BottomScript = str
}
func (this *ViewModelBase) SetCss(str string) {
	this.Css = str
}
