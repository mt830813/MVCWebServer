package Command

type ICommand interface {
	DoCommand(params string)
	GetHelp() string
	GetName() string
}
