package commandhandler

import (
	"github.com/serrano90/cqrs"
	"github.com/serrano90/cqrs/examples/simplecqrs/command"
)

func NewUserCommandHandler() cqrs.CommandHandler {
	c := UserCommandHandler{}
	return &c
}

type UserCommandHandler struct{}

func (handler *UserCommandHandler) Execute(cmd cqrs.Command) (interface{}, error) {
	c := cmd.(*command.UserCommand)
	return c.TypeOf(), nil
}
