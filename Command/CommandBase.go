package Command

type CommandBase struct {
}

func (this *CommandBase) GetName() string {
	return "Name of this command still to be decide"
}

func (this *CommandBase) GetHelp() string {
	return "there is no help for this command"
}
