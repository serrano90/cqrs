package commandbus

import (
	"errors"

	"github.com/serrano90/cqrs"
)

//NewCommandBusInMemory
func NewInMemoryCommandBus() cqrs.CommandBus {
	return &InMemoryCommandBus{
		handlers: make(map[string]cqrs.CommandHandler, 0),
	}
}

//InMemoryCommandBus
type InMemoryCommandBus struct {
	handlers map[string]cqrs.CommandHandler
}

//AddHandler
func (bus *InMemoryCommandBus) AddHandler(handler cqrs.CommandHandler, commands ...cqrs.Command) error {
	for _, command := range commands {
		typeOf := command.TypeOf()
		if _, ok := bus.handlers[typeOf]; ok {
			return errors.New(cqrs.ErrMessageCommandHandlerDuplicated)
		}
		bus.handlers[typeOf] = handler
	}
	return nil
}

//Handler
func (bus *InMemoryCommandBus) Handler(cmd cqrs.Command) (interface{}, error) {
	typeOf := cmd.TypeOf()
	if handler, ok := bus.handlers[typeOf]; ok {
		return handler.Execute(cmd)
	}
	return nil, errors.New(cqrs.ErrMessageCommandHandlerDoesNotExist)
}
