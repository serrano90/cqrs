package cqrs

//CommandHandler is a interface that all commandhandlers implement
type CommandHandler interface {
	//Execute
	Execute(cmd Command) (interface{}, error)
}

//CommandHandlerFunc is a function can by use with as a commandhandler
type CommandHandlerFunc func(cmd Command) (interface{}, error)

func (f CommandHandlerFunc) Execute(cmd Command) (interface{}, error) {
	return f(cmd)
}
