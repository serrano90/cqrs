package cqrs

//CommandBus
type CommandBus interface {
	//RegisterHnadler
	AddHandler(CommandHandler, ...Command) error
	//Handler
	Handler(Command) (interface{}, error)
}
