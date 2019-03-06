package cqrs_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/serrano90/cqrs"

	"github.com/serrano90/converttype"
)

type TestCommand struct {
	name *string
}

func (this *TestCommand) TypeOf() string {
	return reflect.TypeOf(this).String()
}

func TestAddCommandHandlerMiddleware(t *testing.T) {
	c := &TestCommand{name: converttype.String("Name")}

	handlerMiddleware := cqrs.CommandHandlerMiddleware(func(next cqrs.CommandHandler) cqrs.CommandHandler {
		return cqrs.CommandHandlerFunc(func(cmd cqrs.Command) (interface{}, error) {
			testCommand := cmd.(*TestCommand)
			if *testCommand.name != "Name" {
				return nil, errors.New("Incorrect name")
			}
			testCommand.name = converttype.String("Middleware")
			return next.Execute(testCommand)
		})
	})

	handler := cqrs.CommandHandlerFunc(func(cmd cqrs.Command) (interface{}, error) {
		testCommand := cmd.(*TestCommand)
		if *testCommand.name != "Middleware" {
			return nil, errors.New("Incorrect name")
		}
		return "Success", nil
	})

	bus := cqrs.AddCommandHandlerMiddleware(handler, handlerMiddleware)
	response, err := bus.Execute(c)
	if err != nil || response != "Success" {
		t.Fail()
	}
}
