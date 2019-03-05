package main_test

import (
	"testing"

	"github.com/serrano90/converttype"
	"github.com/serrano90/cqrs"
	"github.com/serrano90/cqrs/commandbus"
	"github.com/serrano90/cqrs/examples/simplecqrs/command"
	"github.com/serrano90/cqrs/examples/simplecqrs/commandhandler"
	"github.com/serrano90/cqrs/middleware"
)

var bus cqrs.CommandBus

func setup() {
	bus = commandbus.NewInMemoryCommandBus()

	//load middleware
	validationMiddleware := middleware.NewValidationMiddleware()

	//load all command handler
	usercommandhandler := cqrs.AddCommandHandlerMiddleware(commandhandler.NewUserCommandHandler(), validationMiddleware)

	//load the command handler in bus
	bus.AddHandler(usercommandhandler, &command.UserCommand{})
}

func TestUserHandlerFail(t *testing.T) {
	setup()
	user := command.NewUserCommand(converttype.String(""))
	_, err := bus.Handler(user)
	if err == nil {
		t.Fail()
	}
}

func TestUserHandlerSuccess(t *testing.T) {
	setup()
	user := command.NewUserCommand(converttype.String("testingName"))
	_, err := bus.Handler(user)
	if err != nil {
		t.Error(err)
	}
}
